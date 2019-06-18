// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

func (vm *VM) drop() {
	stackLenStart := len(vm.ctx.stack)
	var opStk uint64
	if len(vm.ctx.stack) > 0 {
		opStk = vm.ctx.stack[0]
	}

	// The operation we're logging
	vm.ctx.stack = vm.ctx.stack[:len(vm.ctx.stack)-1]

	// Log this operation
	stackLenFinish := len(vm.ctx.stack)
	opLog(vm, 0x1A, "Drop", []string{"program_counter", "stack_top", "stack_length_start", "stack_length_finish"},
		[]interface{}{vm.ctx.pc, opStk, stackLenStart, stackLenFinish})
}

func (vm *VM) selectOp() {
	stackLenStart := len(vm.ctx.stack)
	var opStk uint64
	if len(vm.ctx.stack) > 0 {
		opStk = vm.ctx.stack[0]
	}

	// The operation we're logging
	c := vm.popUint32()
	val2 := vm.popUint64()
	val1 := vm.popUint64()

	cond := c != 0
	var val uint64
	if cond {
		val = val1
	} else {
		val = val2
	}
	vm.pushUint64(val)

	// Log this operation
	stackLenFinish := len(vm.ctx.stack)
	opLog(vm, 0x1B, "Select", []string{"program_counter", "stack_top", "condition", "arg_1", "arg_2", "condition_met", "value", "stack_length_start", "stack_length_finish"},
		[]interface{}{vm.ctx.pc, opStk, c, val1, val2, cond, val, stackLenStart, stackLenFinish})
}
