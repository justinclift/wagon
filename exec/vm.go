// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package exec provides functions for executing WebAssembly bytecode.
package exec

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"math"

	"github.com/go-interpreter/wagon/disasm"
	"github.com/go-interpreter/wagon/exec/internal/compile"
	"github.com/go-interpreter/wagon/wasm"
	ops "github.com/go-interpreter/wagon/wasm/operators"
	"github.com/jackc/pgx"
)

var (
	// ErrMultipleLinearMemories is returned by (*VM).NewVM when the module
	// has more then one entries in the linear memory space.
	ErrMultipleLinearMemories = errors.New("exec: more than one linear memories in module")
	// ErrInvalidArgumentCount is returned by (*VM).ExecCode when an invalid
	// number of arguments to the WebAssembly function are passed to it.
	ErrInvalidArgumentCount = errors.New("exec: invalid number of arguments to function")
)

// InvalidReturnTypeError is returned by (*VM).ExecCode when the module
// specifies an invalid return type value for the executed function.
type InvalidReturnTypeError int8

func (e InvalidReturnTypeError) Error() string {
	return fmt.Sprintf("Function has invalid return value_type: %d", int8(e))
}

// InvalidFunctionIndexError is returned by (*VM).ExecCode when the function
// index provided is invalid.
type InvalidFunctionIndexError int64

func (e InvalidFunctionIndexError) Error() string {
	return fmt.Sprintf("Invalid index to function index space: %d", int64(e))
}

type context struct {
	stack   []uint64
	locals  []uint64
	code    []byte
	asm     []asmBlock
	pc      int64
	curFunc int64
}

// VM is the execution context for executing WebAssembly bytecode.
type VM struct {
	ctx context

	module  *wasm.Module
	globals []uint64
	memory  []byte
	funcs   []function

	funcTable [256]func()

	// RecoverPanic controls whether the `ExecCode` method
	// recovers from a panic and returns it as an error
	// instead.
	// A panic can occur either when executing an invalid VM
	// or encountering an invalid instruction, e.g. `unreachable`.
	RecoverPanic bool

	abort bool // Flag for host functions to terminate execution

	nativeBackend *nativeCompiler

	// PostgreSQL pieces, for Operating Logging
	pg       *pgx.ConnPool
	PgTx     *pgx.Tx
	PgRunNum int
}

// As per the WebAssembly spec: https://github.com/WebAssembly/design/blob/27ac254c854994103c24834a994be16f74f54186/Semantics.md#linear-memory
const wasmPageSize = 65536 // (64 KB)

var endianess = binary.LittleEndian

var opNum int // Simple counter for operation logging

type config struct {
	EnableAOT  bool
	PGConnPool *pgx.ConnPool
	PGDBRun    int
}

// VMOption describes a customization that can be applied to the VM.
type VMOption func(c *config)

// EnableAOT enables ahead-of-time compilation of supported opcodes
// into runs of native instructions, if wagon supports native compilation
// for the current architecture.
func EnableAOT(v bool) VMOption {
	return func(c *config) {
		c.EnableAOT = v
	}
}

// PGConnPool passes a pre-established PostgreSQL connection pool, for
// logging all operations through.
func PGConnPool(p *pgx.ConnPool) VMOption {
	return func(c *config) {
		c.PGConnPool = p
	}
}

// PGDBRun passes the "execution run" number, used to identify all logging
// operations in a given execution run.
func PGDBRun(i int) VMOption {
	return func(c *config) {
		c.PGDBRun = i
	}
}

// NewVM creates a new VM from a given module and options. If the module defines
// a start function, it will be executed.
func NewVM(module *wasm.Module, opts ...VMOption) (*VM, error) {
	var (
		vm      VM
		options config
		err     error
	)
	for _, opt := range opts {
		opt(&options)
	}

	// If a PostgreSQL Connection Pool was passed, set up the needed Operation Logging pieces
	if options.PGConnPool != nil {
		// Set the execution run number
		vm.PgRunNum = options.PGDBRun

		// Begin a PostgreSQL transaction
		// TODO: Find out if pgx.BeginBatch() would be useful here, as opposed to changing this to an in-memory
		//       structure, suitable for using with COPY FROM
		vm.pg = options.PGConnPool
		vm.PgTx, err = vm.pg.Begin()
		if err != nil {
			panic(err)
		}
	}

	if module.Memory != nil && len(module.Memory.Entries) != 0 {
		if len(module.Memory.Entries) > 1 {
			return nil, ErrMultipleLinearMemories
		}
		vm.memory = make([]byte, uint(module.Memory.Entries[0].Limits.Initial)*wasmPageSize)
		copy(vm.memory, module.LinearMemoryIndexSpace[0])
	}

	vm.funcs = make([]function, len(module.FunctionIndexSpace)) // Holds the compiled functions
	vm.globals = make([]uint64, len(module.GlobalIndexSpace))
	vm.newFuncTable()
	vm.module = module

	nNatives := 0
	for i, fn := range module.FunctionIndexSpace {
		// Skip native methods as they need not be
		// disassembled; simply add them at the end
		// of the `funcs` array as is, as specified
		// in the spec. See the "host functions"
		// section of:
		// https://webassembly.github.io/spec/core/exec/modules.html#allocation
		if fn.IsHost() {
			vm.funcs[i] = goFunction{
				typ: fn.Host.Type(),
				val: fn.Host,
			}
			nNatives++
			continue
		}

		disassembly, err := disasm.NewDisassembly(fn, module)
		if err != nil {
			return nil, err
		}

		totalLocalVars := 0
		totalLocalVars += len(fn.Sig.ParamTypes)
		for _, entry := range fn.Body.Locals {
			totalLocalVars += int(entry.Count)
		}
		code, meta := compile.Compile(disassembly.Code)
		vm.funcs[i] = compiledFunction{
			codeMeta:       meta,
			code:           code,
			branchTables:   meta.BranchTables,
			maxDepth:       disassembly.MaxDepth,
			totalLocalVars: totalLocalVars,
			args:           len(fn.Sig.ParamTypes),
			returns:        len(fn.Sig.ReturnTypes) != 0,
		}
	}

	if err := vm.resetGlobals(); err != nil {
		return nil, err
	}

	if module.Start != nil {
		_, err := vm.ExecCode(int64(module.Start.Index))
		if err != nil {
			return nil, err
		}
	}

	if options.EnableAOT {
		supportedBackend, backend := nativeBackend()
		if supportedBackend {
			vm.nativeBackend = backend
			if err := vm.tryNativeCompile(); err != nil {
				return nil, err
			}
		}
	}

	return &vm, nil
}

func (vm *VM) resetGlobals() error {
	for i, global := range vm.module.GlobalIndexSpace {
		val, err := vm.module.ExecInitExpr(global.Init)
		if err != nil {
			return err
		}
		switch v := val.(type) {
		case int32:
			vm.globals[i] = uint64(v)
		case int64:
			vm.globals[i] = uint64(v)
		case float32:
			vm.globals[i] = uint64(math.Float32bits(v))
		case float64:
			vm.globals[i] = uint64(math.Float64bits(v))
		}
	}

	return nil
}

// Memory returns the linear memory space for the VM.
func (vm *VM) Memory() []byte {
	return vm.memory
}

func (vm *VM) pushBool(v bool) {
	if v {
		vm.pushUint64(1)
	} else {
		vm.pushUint64(0)
	}
}

func (vm *VM) fetchBool() bool {
	return vm.fetchInt8() != 0
}

func (vm *VM) fetchInt8() int8 {
	i := int8(vm.ctx.code[vm.ctx.pc])
	vm.ctx.pc++
	return i
}

func (vm *VM) fetchUint32() uint32 {
	z := vm.ctx.code[vm.ctx.pc:]
	v := endianess.Uint32(z)
	// v := endianess.Uint32(vm.ctx.code[vm.ctx.pc:])
	vm.ctx.pc += 4
	return v
}

func (vm *VM) fetchInt32() int32 {
	return int32(vm.fetchUint32())
}

func (vm *VM) fetchFloat32() float32 {
	return math.Float32frombits(vm.fetchUint32())
}

func (vm *VM) fetchUint64() uint64 {
	v := endianess.Uint64(vm.ctx.code[vm.ctx.pc:])
	vm.ctx.pc += 8
	return v
}

func (vm *VM) fetchInt64() int64 {
	return int64(vm.fetchUint64())
}

func (vm *VM) fetchFloat64() float64 {
	return math.Float64frombits(vm.fetchUint64())
}

func (vm *VM) popUint64() uint64 {
	i := vm.ctx.stack[len(vm.ctx.stack)-1]
	vm.ctx.stack = vm.ctx.stack[:len(vm.ctx.stack)-1]
	return i
}

func (vm *VM) popInt64() int64 {
	return int64(vm.popUint64())
}

func (vm *VM) popFloat64() float64 {
	return math.Float64frombits(vm.popUint64())
}

func (vm *VM) popUint32() uint32 {
	return uint32(vm.popUint64())
}

func (vm *VM) popInt32() int32 {
	return int32(vm.popUint32())
}

func (vm *VM) popFloat32() float32 {
	return math.Float32frombits(vm.popUint32())
}

func (vm *VM) pushUint64(i uint64) {
	if debugStackDepth {
		if len(vm.ctx.stack) >= cap(vm.ctx.stack) {
			panic("stack exceeding max depth: " + fmt.Sprintf("len=%d,cap=%d", len(vm.ctx.stack), cap(vm.ctx.stack)))
		}
	}
	vm.ctx.stack = append(vm.ctx.stack, i)
}

func (vm *VM) pushInt64(i int64) {
	vm.pushUint64(uint64(i))
}

func (vm *VM) pushFloat64(f float64) {
	vm.pushUint64(math.Float64bits(f))
}

func (vm *VM) pushUint32(i uint32) {
	vm.pushUint64(uint64(i))
}

func (vm *VM) pushInt32(i int32) {
	vm.pushUint64(uint64(i))
}

func (vm *VM) pushFloat32(f float32) {
	vm.pushUint32(math.Float32bits(f))
}

// ExecCode calls the function with the given index and arguments.
// fnIndex should be a valid index into the function index space of
// the VM's module.
func (vm *VM) ExecCode(fnIndex int64, args ...uint64) (rtrn interface{}, err error) {
	// If used as a library, client code should set vm.RecoverPanic to true
	// in order to have an error returned.
	if vm.RecoverPanic {
		defer func() {
			if r := recover(); r != nil {
				switch e := r.(type) {
				case error:
					err = e
				default:
					err = fmt.Errorf("exec: %v", e)
				}
			}
		}()
	}
	if int(fnIndex) > len(vm.funcs) {
		return nil, InvalidFunctionIndexError(fnIndex)
	}
	if len(vm.module.GetFunction(int(fnIndex)).Sig.ParamTypes) != len(args) {
		return nil, ErrInvalidArgumentCount
	}
	compiled, ok := vm.funcs[fnIndex].(compiledFunction)
	if !ok {
		panic(fmt.Sprintf("exec: function at index %d is not a compiled function", fnIndex))
	}

	depth := compiled.maxDepth + 1
	if cap(vm.ctx.stack) < depth {
		vm.ctx.stack = make([]uint64, 0, depth)
	} else {
		vm.ctx.stack = vm.ctx.stack[:0]
	}

	vm.ctx.locals = make([]uint64, compiled.totalLocalVars)
	vm.ctx.pc = 0
	vm.ctx.code = compiled.code
	vm.ctx.asm = compiled.asm
	vm.ctx.curFunc = fnIndex

	for i, arg := range args {
		vm.ctx.locals[i] = arg
	}

	res := vm.execCode(compiled)
	if compiled.returns {
		rtrnType := vm.module.GetFunction(int(fnIndex)).Sig.ReturnTypes[0]
		switch rtrnType {
		case wasm.ValueTypeI32:
			rtrn = uint32(res)
		case wasm.ValueTypeI64:
			rtrn = uint64(res)
		case wasm.ValueTypeF32:
			rtrn = math.Float32frombits(uint32(res))
		case wasm.ValueTypeF64:
			rtrn = math.Float64frombits(res)
		default:
			return nil, InvalidReturnTypeError(rtrnType)
		}
	}

	// Set up an automatic transaction commit for the opLogging
	defer func() {
		err = vm.PgTx.Commit()
		if err != nil {
			panic(err)
		}
	}()

	return rtrn, nil
}

func (vm *VM) execCode(compiled compiledFunction) uint64 {
outer:
	for int(vm.ctx.pc) < len(vm.ctx.code) && !vm.abort {
		op := vm.ctx.code[vm.ctx.pc]
		vm.ctx.pc++
		switch op {
		case ops.Return:

			// Log this operation
			opLog(vm, op, "Return", []string{"program_counter", "stack_start"}, []interface{}{vm.ctx.pc, vm.ctx.stack})

			break outer
		case compile.OpJmp:
			origPC := vm.ctx.pc
			vm.ctx.pc = vm.fetchInt64()

			// Log this operation
			opLog(vm, op, "Jmp unconditional", []string{"program_counter", "stack_start", "target"},
				[]interface{}{origPC, vm.ctx.stack, vm.ctx.pc})
		case compile.OpJmpZ:
			origPC := vm.ctx.pc
			stackStart := vm.ctx.stack

			// The operation we're logging
			target := vm.fetchInt64()
			cond := vm.popUint32() == 0
			if cond {
				vm.ctx.pc = target
			}

			// Log this operation
			opLog(vm, op, "Jmp if zero", []string{"program_counter", "stack_start", "stack_finish", "condition_met", "target"},
				[]interface{}{origPC, stackStart, vm.ctx.stack, cond, target})
		case compile.OpJmpNz:
			origPC := vm.ctx.pc
			stackStart := vm.ctx.stack

			// The operation we're logging
			target := vm.fetchInt64()
			preserveTop := vm.fetchBool()
			discard := vm.fetchInt64()
			cond := vm.popUint32() != 0
			if cond {
				vm.ctx.pc = target
				var top uint64
				if preserveTop {
					top = vm.ctx.stack[len(vm.ctx.stack)-1]
				}
				vm.ctx.stack = vm.ctx.stack[:len(vm.ctx.stack)-int(discard)]
				if preserveTop {
					vm.pushUint64(top)
				}
			}

			// Log this operation
			opLog(vm, op, "Jmp if Not Zero / branch if", []string{"program_counter", "stack_start", "stack_finish", "target", "preserve_top", "discard", "condition_met"},
				[]interface{}{origPC, stackStart, vm.ctx.stack, target, preserveTop, discard, cond})
		case ops.BrTable:
			index := vm.fetchInt64()
			label := vm.popInt32()
			cf, ok := vm.funcs[vm.ctx.curFunc].(compiledFunction)
			if !ok {
				panic(fmt.Sprintf("exec: function at index %d is not a compiled function", vm.ctx.curFunc))
			}
			table := cf.branchTables[index]
			var target compile.Target
			if label >= 0 && label < int32(len(table.Targets)) {
				target = table.Targets[int32(label)]
			} else {
				target = table.DefaultTarget
			}

			if target.Return {
				break outer
			}
			vm.ctx.pc = target.Addr
			var top uint64
			if target.PreserveTop {
				top = vm.ctx.stack[len(vm.ctx.stack)-1]
			}
			vm.ctx.stack = vm.ctx.stack[:len(vm.ctx.stack)-int(target.Discard)]
			if target.PreserveTop {
				vm.pushUint64(top)
			}
			continue
		case compile.OpDiscard:
			stackStart := append(make([]uint64, len(vm.ctx.stack)), vm.ctx.stack...) // Create a separate copy, to be safe

			// The operation we're logging
			place := vm.fetchInt64()
			vm.ctx.stack = vm.ctx.stack[:len(vm.ctx.stack)-int(place)]

			// Log this operation
			opLog(vm, op, "Discard", []string{"program_counter", "stack_start", "stack_finish"},
				[]interface{}{vm.ctx.pc, stackStart, vm.ctx.stack})
		case compile.OpDiscardPreserveTop:
			stackStart := append(make([]uint64, len(vm.ctx.stack)), vm.ctx.stack...) // Create a separate copy, to be safe

			// The operation we're logging
			top := vm.ctx.stack[len(vm.ctx.stack)-1]
			place := vm.fetchInt64()
			vm.ctx.stack = vm.ctx.stack[:len(vm.ctx.stack)-int(place)]
			vm.pushUint64(top)

			// Log this operation
			opLog(vm, op, "Discard preserving top stack value", []string{"program_counter", "stack_start", "stack_finish"},
				[]interface{}{vm.ctx.pc, stackStart, vm.ctx.stack})
		case ops.WagonNativeExec:
			// Log this operation
			opLog(vm, op, "Wagon native execution op - shouldn't happen", []string{"program_counter", "stack_start"},
				[]interface{}{vm.ctx.pc, vm.ctx.stack})

			// The operation we're logging
			i := vm.fetchUint32()
			vm.nativeCodeInvocation(i)
		default:
			vm.funcTable[op]()
		}
	}

	if compiled.returns {
		return vm.ctx.stack[len(vm.ctx.stack)-1]
	}
	return 0
}

// Restart readies the VM for another run.
func (vm *VM) Restart() {
	vm.resetGlobals()
	vm.ctx.locals = make([]uint64, 0)
	vm.abort = false
}

// Close frees any resources managed by the VM.
func (vm *VM) Close() error {
	vm.abort = true // prevents further use.
	if vm.nativeBackend != nil {
		if err := vm.nativeBackend.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Process is a proxy passed to host functions in order to access
// things such as memory and control.
type Process struct {
	vm *VM
}

// NewProcess creates a VM interface object for host functions
func NewProcess(vm *VM) *Process {
	return &Process{vm: vm}
}

// ReadAt implements the ReaderAt interface: it copies into p
// the content of memory at offset off.
func (proc *Process) ReadAt(p []byte, off int64) (int, error) {
	mem := proc.vm.Memory()

	var length int
	if len(mem) < len(p)+int(off) {
		length = len(mem) - int(off)
	} else {
		length = len(p)
	}

	copy(p, mem[off:off+int64(length)])

	var err error
	if length < len(p) {
		err = io.ErrShortBuffer
	}

	return length, err
}

// WriteAt implements the WriterAt interface: it writes the content of p
// into the VM memory at offset off.
func (proc *Process) WriteAt(p []byte, off int64) (int, error) {
	mem := proc.vm.Memory()

	var length int
	if len(mem) < len(p)+int(off) {
		length = len(mem) - int(off)
	} else {
		length = len(p)
	}

	copy(mem[off:], p[:length])

	var err error
	if length < len(p) {
		err = io.ErrShortWrite
	}

	return length, err
}

// Terminate stops the execution of the current module.
func (proc *Process) Terminate() {
	proc.vm.abort = true
}

// Send the opcode data to the database for post-run analysis.  For now we don't return any error code, just to keep
// the likely bulk code changes somewhat simple
func opLog(vm *VM, opCode byte, opName string, fields []string, data []interface{}) {
	if vm.pg == nil {
		// Operating logging isn't enabled
		return
	}
	if len(fields) != len(data) {
		log.Print("Mismatching field and data count to opLog()")
		return
	}
	var s, t string
	for i, j := range fields {
		s += ", " + j
		t += fmt.Sprintf(", $%d", 5+i)
	}
	dbQuery := fmt.Sprintf(`
		INSERT INTO execution_run (op_num, run_num, op_code, op_name%s)
		VALUES ($1, $2, $3, $4%s)`, s, t)
	var err error
	var commandTag pgx.CommandTag
	// TODO: Surely there's a better way than this?
	switch len(fields) {
	case 0:
		commandTag, err = vm.PgTx.Exec(dbQuery, opNum, vm.PgRunNum, opCode, opName)
	case 1:
		commandTag, err = vm.PgTx.Exec(dbQuery, opNum, vm.PgRunNum, opCode, opName, data[0])
	case 2:
		commandTag, err = vm.PgTx.Exec(dbQuery, opNum, vm.PgRunNum, opCode, opName, data[0], data[1])
	case 3:
		commandTag, err = vm.PgTx.Exec(dbQuery, opNum, vm.PgRunNum, opCode, opName, data[0], data[1], data[2])
	case 4:
		commandTag, err = vm.PgTx.Exec(dbQuery, opNum, vm.PgRunNum, opCode, opName, data[0], data[1], data[2], data[3])
	case 5:
		commandTag, err = vm.PgTx.Exec(dbQuery, opNum, vm.PgRunNum, opCode, opName, data[0], data[1], data[2], data[3], data[4])
	case 6:
		commandTag, err = vm.PgTx.Exec(dbQuery, opNum, vm.PgRunNum, opCode, opName, data[0], data[1], data[2], data[3], data[4], data[5])
	case 7:
		commandTag, err = vm.PgTx.Exec(dbQuery, opNum, vm.PgRunNum, opCode, opName, data[0], data[1], data[2], data[3], data[4], data[5], data[6])
	case 8:
		commandTag, err = vm.PgTx.Exec(dbQuery, opNum, vm.PgRunNum, opCode, opName, data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7])
	default:
		log.Printf("Need to add a case for %d to the opLog() function", len(fields))
		return
	}
	if err != nil {
		log.Print(err)
		return
	}
	if numRows := commandTag.RowsAffected(); numRows != 1 {
		log.Printf("Wrong number of rows (%v) affected when logging an operation: %v\n", numRows, opName)
	}

	// Commit every 10k inserts, so quitting via Ctrl+C keeps the majority of info thus far
	if (opNum % 10000) == 0 {
		err = vm.PgTx.Commit()
		if err != nil {
			panic(err)
		}
		vm.PgTx, err = vm.pg.Begin()
		if err != nil {
			panic(err)
		}
	}
	opNum++
	return
}
