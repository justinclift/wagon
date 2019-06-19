// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

func (vm *VM) getLocal() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	index := vm.fetchUint32()
	val := vm.ctx.locals[int(index)]
	vm.pushUint64(val)

	// Log this operation
	opLog(vm, 0x20, "Get local", []string{"program_counter", "local_id", "value", "locals_start", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, index, val, vm.ctx.locals, stackStart, vm.ctx.stack})
}

func (vm *VM) setLocal() {
	stackStart := vm.ctx.stack
	var localsStart []uint64
	localsStart = append(localsStart, vm.ctx.locals...)

	// The operation we're logging
	index := vm.fetchUint32()
	val := vm.popUint64()
	vm.ctx.locals[int(index)] = val

	// Log this operation
	opLog(vm, 0x21, "Set local", []string{"program_counter", "local_id", "value", "locals_start", "locals_finish", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, index, val, localsStart, vm.ctx.locals, stackStart, vm.ctx.stack})
}

func (vm *VM) teeLocal() {
	stackStart := vm.ctx.stack
	localsStart := make([]uint64, len(vm.ctx.locals))
	copy(localsStart, vm.ctx.locals)

	// The operation we're logging
	index := vm.fetchUint32()
	val := vm.ctx.stack[len(vm.ctx.stack)-1]
	vm.ctx.locals[int(index)] = val

	// Log this operation
	opLog(vm, 0x22, "Tee local", []string{"program_counter", "local_id", "value", "locals_start", "locals_finish", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, index, val, localsStart, vm.ctx.locals, stackStart, vm.ctx.stack})
}

func (vm *VM) getGlobal() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	index := vm.fetchUint32()
	val := vm.globals[int(index)]
	vm.pushUint64(val)

	// Log this operation
	opLog(vm, 0x23, "Get global", []string{"program_counter", "from_global", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, index, val, stackStart, vm.ctx.stack})
}

func (vm *VM) setGlobal() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	index := vm.fetchUint32()
	val := vm.popUint64()
	vm.globals[int(index)] = val

	// Log this operation
	opLog(vm, 0x24, "Set global", []string{"program_counter", "to_global", "value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, index, val, stackStart, vm.ctx.stack})
}
