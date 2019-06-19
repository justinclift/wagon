// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

func (vm *VM) i32Const() {
	stackStart := vm.ctx.stack

	z := vm.fetchUint32()
	vm.pushUint32(z)

	stackFinish := vm.ctx.stack
	opLog(vm, 0x41, "i32 constant", []string{"program_counter", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, z, stackStart, stackFinish})
}

func (vm *VM) i64Const() {
	stackStart := vm.ctx.stack

	z := vm.fetchUint64()
	vm.pushUint64(z)

	opLog(vm, 0x42, "i64 constant", []string{"program_counter", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, z, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Const() {
	stackStart := vm.ctx.stack

	z := vm.fetchFloat32()
	vm.pushFloat32(z)

	opLog(vm, 0x43, "f32 constant", []string{"program_counter", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, z, stackStart, vm.ctx.stack})
}

func (vm *VM) f64Const() {
	stackStart := vm.ctx.stack

	z := vm.fetchFloat64()
	vm.pushFloat64(z)

	opLog(vm, 0x44, "f64 constant", []string{"program_counter", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, z, stackStart, vm.ctx.stack})
}
