// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import (
	"math"
)

func (vm *VM) i32Wrapi64() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popUint64()
	val := uint32(v1)
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0xA7, "i32 Wrap i64", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32TruncSF32() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat32()
	val := int32(math.Trunc(float64(v1)))
	vm.pushInt32(val)

	// Log this operation
	opLog(vm, 0xA8, "i32 Truncate f32 signed", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32TruncUF32() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat32()
	val := uint32(math.Trunc(float64(v1)))
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0xA9, "i32 Truncate f32 unsigned", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32TruncSF64() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat64()
	val := int32(math.Trunc(v1))
	vm.pushInt32(val)

	// Log this operation
	opLog(vm, 0xAA, "i32 Truncate f64 signed", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32TruncUF64() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat64()
	val := uint32(math.Trunc(v1))
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0xAB, "i32 Truncate f64 unsigned", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i64ExtendSI32() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popInt32()
	val := int64(v1)
	vm.pushInt64(val)

	// Log this operation
	opLog(vm, 0xAC, "i64 Extend i32 signed", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i64ExtendUI32() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popUint32()
	val := uint64(v1)
	vm.pushUint64(val)

	// Log this operation
	opLog(vm, 0xAD, "i64 Extend i32 Unsigned", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i64TruncSF32() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat32()
	val := int64(math.Trunc(float64(v1)))
	vm.pushInt64(val)

	// Log this operation
	opLog(vm, 0xAE, "i64 Truncate f32 signed", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i64TruncUF32() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat32()
	val := uint64(math.Trunc(float64(v1)))
	vm.pushUint64(val)

	// Log this operation
	opLog(vm, 0xAF, "i64 Truncate f32 unsigned", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i64TruncSF64() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat64()
	val := int64(math.Trunc(v1))
	vm.pushInt64(val)

	// Log this operation
	opLog(vm, 0xB0, "i64 Truncate f64 signed", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i64TruncUF64() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat64()
	val := uint64(math.Trunc(v1))
	vm.pushUint64(val)

	// Log this operation
	opLog(vm, 0xB1, "i64 Truncate f64 unsigned", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32ConvertSI32() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popInt32()
	val := float32(v1)
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0xB2, "f32 Convert i32 signed", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32ConvertUI32() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popUint32()
	val := float32(v1)
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0xB3, "f32 Convert i32 unsigned", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32ConvertSI64() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popInt64()
	val := float32(v1)
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0xB4, "f32 Convert i64 signed", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32ConvertUI64() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popUint64()
	val := float32(v1)
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0xB5, "f32 Convert i64 unsigned", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32DemoteF64() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat64()
	val := float32(v1)
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0xB6, "f32 Demote f64", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f64ConvertSI32() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popInt32()
	val := float64(v1)
	vm.pushFloat64(val)

	// Log this operation
	opLog(vm, 0xB7, "f64 Convert i32 signed", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f64ConvertUI32() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popUint32()
	val := float64(v1)
	vm.pushFloat64(val)

	// Log this operation
	opLog(vm, 0xB8, "f64 Convert i32 unsigned", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f64ConvertSI64() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popInt64()
	val := float64(v1)
	vm.pushFloat64(val)

	// Log this operation
	opLog(vm, 0xB9, "f64 Convert i64 signed", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f64ConvertUI64() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popUint64()
	val := float64(v1)
	vm.pushFloat64(val)

	// Log this operation
	opLog(vm, 0xBA, "f64 Convert i64 unsigned", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f64PromoteF32() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat32()
	val := float64(v1)
	vm.pushFloat64(val)

	// Log this operation
	opLog(vm, 0xBB, "f64 Promote f32", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}
