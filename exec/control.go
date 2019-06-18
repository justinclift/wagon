// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import "errors"

// ErrUnreachable is the error value used while trapping the VM when
// an unreachable operator is reached during execution.
var ErrUnreachable = errors.New("exec: reached unreachable")

func (vm *VM) unreachable() {
	panic(ErrUnreachable)
}

func (vm *VM) nop() {
	// Log this operation
	var opStk uint64
	if len(vm.ctx.stack) > 0 {
		opStk = vm.ctx.stack[0]
	}
	opLog(vm, 0x1, "Nop", []string{"program_counter", "stack_top"},
		[]interface{}{vm.ctx.pc, opStk})
}
