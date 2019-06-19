// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import (
	"errors"
	"math"
)

// ErrOutOfBoundsMemoryAccess is the error value used while trapping the VM
// when it detects an out of bounds access to the linear memory.
var ErrOutOfBoundsMemoryAccess = errors.New("exec: out of bounds memory access")

func (vm *VM) fetchBaseAddr() int {
	return int(vm.fetchUint32() + uint32(vm.popInt32()))
}

// inBounds returns true when the next vm.fetchBaseAddr() + offset
// indices are in bounds accesses to the linear memory.
func (vm *VM) inBounds(offset int) bool {
	addr := endianess.Uint32(vm.ctx.code[vm.ctx.pc:]) + uint32(vm.ctx.stack[len(vm.ctx.stack)-1])
	return int(addr)+offset < len(vm.memory)
}

// curMem returns a slice to the memory segment pointed to by
// the current base address on the bytecode stream.
func (vm *VM) curMem() []byte {
	return vm.memory[vm.fetchBaseAddr():]
}

func (vm *VM) i32Load() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	addr := vm.fetchBaseAddr()
	val := endianess.Uint32(vm.memory[addr:])
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x28, "i32 load", []string{"program_counter", "memory_address", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, addr, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Load8s() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	if !vm.inBounds(0) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	addr := vm.fetchBaseAddr()
	val := int32(int8(vm.memory[addr]))
	vm.pushInt32(val)

	// Log this operation
	opLog(vm, 0x2C, "i32 load 8-bit signed", []string{"program_counter", "memory_address", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, addr, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Load8u() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	if !vm.inBounds(0) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	addr := vm.fetchBaseAddr()
	val := uint32(uint8(vm.memory[addr]))
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x2D, "i32 load 8-bit unsigned", []string{"program_counter", "memory_address", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, addr, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Load16s() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	if !vm.inBounds(1) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	addr := vm.curMem()
	val := int32(int16(endianess.Uint16(addr)))
	vm.pushInt32(val)

	// Log this operation
	opLog(vm, 0x2E, "i32 load 16-bit signed", []string{"program_counter", "memory_address", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, addr, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Load16u() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	if !vm.inBounds(1) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	addr := vm.curMem()
	val := uint32(endianess.Uint16(addr))
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x2F, "i32 load 16-bit unsigned", []string{"program_counter", "memory_address", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, addr, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i64Load() {
	if !vm.inBounds(7) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushUint64(endianess.Uint64(vm.curMem()))
}

func (vm *VM) i64Load8s() {
	if !vm.inBounds(0) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushInt64(int64(int8(vm.memory[vm.fetchBaseAddr()])))
}

func (vm *VM) i64Load8u() {
	if !vm.inBounds(0) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushUint64(uint64(uint8(vm.memory[vm.fetchBaseAddr()])))
}

func (vm *VM) i64Load16s() {
	if !vm.inBounds(1) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushInt64(int64(int16(endianess.Uint16(vm.curMem()))))
}

func (vm *VM) i64Load16u() {
	if !vm.inBounds(1) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushUint64(uint64(endianess.Uint16(vm.curMem())))
}

func (vm *VM) i64Load32s() {
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushInt64(int64(int32(endianess.Uint32(vm.curMem()))))
}

func (vm *VM) i64Load32u() {
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushUint64(uint64(endianess.Uint32(vm.curMem())))
}

func (vm *VM) f32Store() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	val := math.Float32bits(vm.popFloat32())
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	addr := vm.fetchBaseAddr()
	mem := vm.memory[addr:]
	endianess.PutUint32(mem, val)

	// Log this operation
	opLog(vm, 0x38, "f32 store", []string{"program_counter", "memory_address", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, addr, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Load() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	addr := vm.fetchBaseAddr()
	mem := vm.memory[addr:]
	val := math.Float32frombits(endianess.Uint32(mem))
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x2A, "f32 load", []string{"program_counter", "memory_address", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, addr, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f64Store() {
	v := math.Float64bits(vm.popFloat64())
	if !vm.inBounds(7) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	endianess.PutUint64(vm.curMem(), v)
}

func (vm *VM) f64Load() {
	if !vm.inBounds(7) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.pushFloat64(math.Float64frombits(endianess.Uint64(vm.curMem())))
}

func (vm *VM) i32Store() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	val := vm.popUint32()
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	addr := vm.fetchBaseAddr()
	mem := vm.memory[addr:]
	endianess.PutUint32(mem, val)

	// Log this operation
	opLog(vm, 0x36, "i32 store", []string{"program_counter", "memory_address", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, addr, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Store8() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	val := byte(uint8(vm.popUint32()))
	if !vm.inBounds(0) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	addr := vm.fetchBaseAddr()
	vm.memory[addr] = val

	// Log this operation
	opLog(vm, 0x3A, "i32 store 8-bit", []string{"program_counter", "memory_address", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, addr, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Store16() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	val := uint16(vm.popUint32())
	if !vm.inBounds(1) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	addr := vm.fetchBaseAddr()
	mem := vm.memory[addr:]
	endianess.PutUint16(mem, val)

	// Log this operation
	opLog(vm, 0x3B, "i32 store 16-bit", []string{"program_counter", "memory_address", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, addr, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i64Store() {
	v := vm.popUint64()
	if !vm.inBounds(7) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	endianess.PutUint64(vm.curMem(), v)
}

func (vm *VM) i64Store8() {
	v := byte(uint8(vm.popUint64()))
	if !vm.inBounds(0) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	vm.memory[vm.fetchBaseAddr()] = v
}

func (vm *VM) i64Store16() {
	v := uint16(vm.popUint64())
	if !vm.inBounds(1) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	endianess.PutUint16(vm.curMem(), v)
}

func (vm *VM) i64Store32() {
	v := uint32(vm.popUint64())
	if !vm.inBounds(3) {
		panic(ErrOutOfBoundsMemoryAccess)
	}
	endianess.PutUint32(vm.curMem(), v)
}

func (vm *VM) currentMemory() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	_ = vm.fetchInt8() // reserved (https://github.com/WebAssembly/design/blob/27ac254c854994103c24834a994be16f74f54186/BinaryEncoding.md#memory-related-operators-described-here)
	val := int32(len(vm.memory) / wasmPageSize)
	vm.pushInt32(val)

	// Log this operation
	opLog(vm, 0x3F, "current memory size", []string{"program_counter", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, val, stackStart, vm.ctx.stack})
}

func (vm *VM) growMemory() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	_ = vm.fetchInt8() // reserved (https://github.com/WebAssembly/design/blob/27ac254c854994103c24834a994be16f74f54186/BinaryEncoding.md#memory-related-operators-described-here)
	curLen := len(vm.memory) / wasmPageSize
	n := vm.popInt32()
	vm.memory = append(vm.memory, make([]byte, n*wasmPageSize)...)
	vm.pushInt32(int32(curLen))

	// Log this operation
	opLog(vm, 0x40, "grow memory", []string{"program_counter", "modifier_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, n, stackStart, vm.ctx.stack})
}
