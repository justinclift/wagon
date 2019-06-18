// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

func (vm *VM) getLocal() {
	stackLenStart := len(vm.ctx.stack)
	var opStk uint64
	if len(vm.ctx.stack) > 0 {
		opStk = vm.ctx.stack[0]
	}

	// The operation we're logging
	index := vm.fetchUint32()
	val := vm.ctx.locals[int(index)]
	vm.pushUint64(val)

	// Log this operation
	stackLenFinish := len(vm.ctx.stack)
	opLog(vm, 0x20, "Get local", []string{"program_counter", "stack_top", "local_id", "value", "stack_length_start", "stack_length_finish"},
		[]interface{}{vm.ctx.pc, opStk, index, val, stackLenStart, stackLenFinish})
}

func (vm *VM) setLocal() {
	stackLenStart := len(vm.ctx.stack)
	var opStk uint64
	if len(vm.ctx.stack) > 0 {
		opStk = vm.ctx.stack[0]
	}

	// The operation we're logging
	index := vm.fetchUint32()
	val := vm.popUint64()
	vm.ctx.locals[int(index)] = val

	// Log this operation
	stackLenFinish := len(vm.ctx.stack)
	opLog(vm, 0x21, "Set local", []string{"program_counter", "stack_top", "local_id", "value", "stack_length_start", "stack_length_finish"},
		[]interface{}{vm.ctx.pc, opStk, index, val, stackLenStart, stackLenFinish})
}

func (vm *VM) teeLocal() {
	stackLenStart := len(vm.ctx.stack)
	var opStk uint64
	if len(vm.ctx.stack) > 0 {
		opStk = vm.ctx.stack[0]
	}

	// The operation we're logging
	index := vm.fetchUint32()
	val := vm.ctx.stack[len(vm.ctx.stack)-1]
	vm.ctx.locals[int(index)] = val

	// Log this operation
	stackLenFinish := len(vm.ctx.stack)
	opLog(vm, 0x22, "Tee local", []string{"program_counter", "stack_top", "local_id", "value", "stack_length_start", "stack_length_finish"},
		[]interface{}{vm.ctx.pc, opStk, index, val, stackLenStart, stackLenFinish})
}

func (vm *VM) getGlobal() {
	stackLenStart := len(vm.ctx.stack)
	var opStk uint64
	if len(vm.ctx.stack) > 0 {
		opStk = vm.ctx.stack[0]
	}

	// The operation we're logging
	index := vm.fetchUint32()
	val := vm.globals[int(index)]
	vm.pushUint64(val)

	// Log this operation
	stackLenFinish := len(vm.ctx.stack)
	opLog(vm, 0x23, "Get global", []string{"program_counter", "stack_top", "from_global", "value", "stack_length_start", "stack_length_finish"},
		[]interface{}{vm.ctx.pc, opStk, index, val, stackLenStart, stackLenFinish})
}

func (vm *VM) setGlobal() {
	stackLenStart := len(vm.ctx.stack)
	var opStk uint64
	if len(vm.ctx.stack) > 0 {
		opStk = vm.ctx.stack[0]
	}

	// The operation we're logging
	index := vm.fetchUint32()
	val := vm.popUint64()
	vm.globals[int(index)] = val

	// Log this operation
	stackLenFinish := len(vm.ctx.stack)
	opLog(vm, 0x24, "Set global", []string{"program_counter", "stack_top", "to_global", "value", "stack_length_start", "stack_length_finish"},
		[]interface{}{vm.ctx.pc, opStk, index, val, stackLenStart, stackLenFinish})
}
