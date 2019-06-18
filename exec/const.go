// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

func (vm *VM) i32Const() {
	stackLenStart := len(vm.ctx.stack)

	z := vm.fetchUint32()
	vm.pushUint32(z)

	stackLenFinish := len(vm.ctx.stack)
	opStk := vm.ctx.stack[0]
	opLog(vm, 0x41, "i32 constant", []string{"program_counter", "value", "stack_top", "stack_length_start", "stack_length_finish"},
		[]interface{}{vm.ctx.pc, z, opStk, stackLenStart, stackLenFinish})
}

func (vm *VM) i64Const() {
	stackLenStart := len(vm.ctx.stack)

	z := vm.fetchUint64()
	vm.pushUint64(z)

	stackLenFinish := len(vm.ctx.stack)
	opStk := vm.ctx.stack[0]
	opLog(vm, 0x42, "i64 constant", []string{"program_counter", "value", "stack_top", "stack_length_start", "stack_length_finish"},
		[]interface{}{vm.ctx.pc, z, opStk, stackLenStart, stackLenFinish})
}

func (vm *VM) f32Const() {
	stackLenStart := len(vm.ctx.stack)

	z := vm.fetchFloat32()
	vm.pushFloat32(z)

	stackLenFinish := len(vm.ctx.stack)
	opStk := vm.ctx.stack[0]
	opLog(vm, 0x43, "f32 constant", []string{"program_counter", "value", "stack_top", "stack_length_start", "stack_length_finish"},
		[]interface{}{vm.ctx.pc, z, opStk, stackLenStart, stackLenFinish})
}

func (vm *VM) f64Const() {
	stackLenStart := len(vm.ctx.stack)

	z := vm.fetchFloat64()
	vm.pushFloat64(z)

	stackLenFinish := len(vm.ctx.stack)
	opStk := vm.ctx.stack[0]
	opLog(vm, 0x44, "f64 constant", []string{"program_counter", "value", "stack_top", "stack_length_start", "stack_length_finish"},
		[]interface{}{vm.ctx.pc, z, opStk, stackLenStart, stackLenFinish})
}
