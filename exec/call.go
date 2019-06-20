// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import "errors"

var (
	// ErrSignatureMismatch is the error value used while trapping the VM when
	// a signature mismatch between the table entry and the type entry is found
	// in a call_indirect operation.
	ErrSignatureMismatch = errors.New("exec: signature mismatch in call_indirect")
	// ErrUndefinedElementIndex is the error value used while trapping the VM when
	// an invalid index to the module's table space is used as an operand to
	// call_indirect
	ErrUndefinedElementIndex = errors.New("exec: undefined element index")
)

func (vm *VM) call() {
	stackStart := vm.ctx.stack

	// Fetch the number of the function to call
	index := vm.fetchUint32()


	// Log the start of this operation
	fName := vm.module.FunctionIndexSpace[index].Name
	opLog(vm, 0x10, "Call function start", []string{"program_counter", "function_id", "function_name", "stack_start"},
		[]interface{}{vm.ctx.pc, index, fName, stackStart})

	// Do the call
	vm.funcs[index].call(vm, int64(index))

	// Log the end of this operation
	opLog(vm, 0x10, "Call function end", []string{"program_counter", "function_id", "function_name", "stack_finish"},
		[]interface{}{vm.ctx.pc, index, fName, vm.ctx.stack})
}

func (vm *VM) callIndirect() {
	stackStart := vm.ctx.stack

	index := vm.fetchUint32()
	fnExpect := vm.module.Types.Entries[index]
	_ = vm.fetchUint32() // reserved (https://github.com/WebAssembly/design/blob/27ac254c854994103c24834a994be16f74f54186/BinaryEncoding.md#call-operators-described-here)
	tableIndex := vm.popUint32()
	if int(tableIndex) >= len(vm.module.TableIndexSpace[0]) {
		panic(ErrUndefinedElementIndex)
	}
	elemIndex := vm.module.TableIndexSpace[0][tableIndex]
	fnActual := vm.module.FunctionIndexSpace[elemIndex]

	if len(fnExpect.ParamTypes) != len(fnActual.Sig.ParamTypes) {
		panic(ErrSignatureMismatch)
	}
	if len(fnExpect.ReturnTypes) != len(fnActual.Sig.ReturnTypes) {
		panic(ErrSignatureMismatch)
	}

	for i := range fnExpect.ParamTypes {
		if fnExpect.ParamTypes[i] != fnActual.Sig.ParamTypes[i] {
			panic(ErrSignatureMismatch)
		}
	}

	for i := range fnExpect.ReturnTypes {
		if fnExpect.ReturnTypes[i] != fnActual.Sig.ReturnTypes[i] {
			panic(ErrSignatureMismatch)
		}
	}

	// Log the start of this operation
	opLog(vm, 0x11, "Call indirect function start", []string{"program_counter", "function_id", "stack_start"},
		[]interface{}{vm.ctx.pc, index, stackStart})

	vm.funcs[elemIndex].call(vm, int64(elemIndex))

	// Log the end of this operation
	opLog(vm, 0x11, "Call indirect function end", []string{"program_counter", "function_id", "stack_finish"},
		[]interface{}{vm.ctx.pc, index, vm.ctx.stack})
}
