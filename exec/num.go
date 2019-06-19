// Copyright 2017 The go-interpreter Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import (
	"math"
	"math/bits"
)

// int32 operators

func (vm *VM) i32Clz() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popUint32()
	val := uint64(bits.LeadingZeros32(v1))
	vm.pushUint64(val)

	// Log this operation
	opLog(vm, 0x67, "i32 Count leading zero bits", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Ctz() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popUint32()
	val := uint64(bits.TrailingZeros32(v1))
	vm.pushUint64(val)

	// Log this operation
	opLog(vm, 0x68, "i32 Count trailing zero bits", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Popcnt() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popUint32()
	val := uint64(bits.OnesCount32(v1))
	vm.pushUint64(val)

	// Log this operation
	opLog(vm, 0x69, "i32 Count number of one bits", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Add() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popUint32()
	v2 := vm.popUint32()
	val := v1 + v2
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x6A, "i32 Add", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Sub() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	val := v1 - v2
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x6B, "i32 Sub", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Mul() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	val := v1 * v2
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x6C, "i32 Multiply", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32DivS() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popInt32()
	v1 := vm.popInt32()
	val := v1 / v2
	vm.pushInt32(val)

	// Log this operation
	opLog(vm, 0x6D, "i32 Divide signed", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32DivU() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	val := v1 / v2
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x6E, "i32 Divide unsigned", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32RemS() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popInt32()
	v1 := vm.popInt32()
	val := v1 % v2
	vm.pushInt32(val)

	// Log this operation
	opLog(vm, 0x6F, "i32 Remainder signed", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32RemU() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	val := v1 % v2
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x70, "i32 Remainder unsigned", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32And() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	val := v1 & v2
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x71, "i32 And", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Or() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	val := v1 | v2
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x72, "i32 Or", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Xor() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	val := v1 ^ v2
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x73, "i32 Xor", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Shl() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	val := v1 << v2
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x74, "i32 Shift left", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32ShrS() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popInt32()
	val := v1 >> v2
	vm.pushInt32(val)

	// Log this operation
	opLog(vm, 0x75, "i32 Shift right signed", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32ShrU() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	val := v1 >> v2
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x76, "i32 Shift right unsigned", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Rotl() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	val := bits.RotateLeft32(v1, int(v2))
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x77, "i32 Rotate left", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Rotr() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	val := bits.RotateLeft32(v1, -int(v2))
	vm.pushUint32(val)

	// Log this operation
	opLog(vm, 0x78, "i32 Rotate right", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) i32LeS() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popInt32()
	v1 := vm.popInt32()
	cond := v1 <= v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x4C, "i32 Less than or equal signed", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) i32LeU() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	cond := v1 <= v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x4D, "i32 Less than or equal unsigned", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) i32LtS() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popInt32()
	v1 := vm.popInt32()
	cond := v1 < v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x48, "i32 Less than signed", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) i32LtU() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	cond := v1 < v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x49, "i32 Less than unsigned", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) i32GtS() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popInt32()
	v1 := vm.popInt32()
	cond := v1 > v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x4A, "i32 Greater than signed", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) i32GtU() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	cond := v1 > v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x4B, "i32 Greater than unsigned", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) i32GeS() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popInt32()
	v1 := vm.popInt32()
	cond := v1 >= v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x4E, "i32 Greater than or equal signed", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) i32GeU() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	cond := v1 >= v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x4F, "i32 Greater than or equal unsigned", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Eqz() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	val := vm.popUint32()
	cond := val == 0
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x45, "i32 Equal to zero", []string{"program_counter", "value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, val, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Eq() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	arg1 := vm.popUint32()
	arg2 := vm.popUint32()
	cond := arg1 == arg2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x46, "i32 Equal", []string{"program_counter", "arg_1", "arg_2", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, arg1, arg2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) i32Ne() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popUint32()
	v1 := vm.popUint32()
	cond := v1 != v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x47, "i32 Not equal", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

// int64 operators

func (vm *VM) i64Clz() {
	vm.pushUint64(uint64(bits.LeadingZeros64(vm.popUint64())))
}

func (vm *VM) i64Ctz() {
	vm.pushUint64(uint64(bits.TrailingZeros64(vm.popUint64())))
}

func (vm *VM) i64Popcnt() {
	vm.pushUint64(uint64(bits.OnesCount64(vm.popUint64())))
}

func (vm *VM) i64Add() {
	vm.pushUint64(vm.popUint64() + vm.popUint64())
}

func (vm *VM) i64Sub() {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushUint64(v1 - v2)
}

func (vm *VM) i64Mul() {
	vm.pushUint64(vm.popUint64() * vm.popUint64())
}

func (vm *VM) i64DivS() {
	v2 := vm.popInt64()
	v1 := vm.popInt64()
	vm.pushInt64(v1 / v2)
}

func (vm *VM) i64DivU() {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushUint64(v1 / v2)
}

func (vm *VM) i64RemS() {
	v2 := vm.popInt64()
	v1 := vm.popInt64()
	vm.pushInt64(v1 % v2)
}

func (vm *VM) i64RemU() {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushUint64(v1 % v2)
}

func (vm *VM) i64And() {
	vm.pushUint64(vm.popUint64() & vm.popUint64())
}

func (vm *VM) i64Or() {
	vm.pushUint64(vm.popUint64() | vm.popUint64())
}

func (vm *VM) i64Xor() {
	vm.pushUint64(vm.popUint64() ^ vm.popUint64())
}

func (vm *VM) i64Shl() {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushUint64(v1 << v2)
}

func (vm *VM) i64ShrS() {
	v2 := vm.popUint64()
	v1 := vm.popInt64()
	vm.pushInt64(v1 >> v2)
}

func (vm *VM) i64ShrU() {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushUint64(v1 >> v2)
}

func (vm *VM) i64Rotl() {
	v2 := vm.popInt64()
	v1 := vm.popUint64()
	vm.pushUint64(bits.RotateLeft64(v1, int(v2)))
}

func (vm *VM) i64Rotr() {
	v2 := vm.popInt64()
	v1 := vm.popUint64()
	vm.pushUint64(bits.RotateLeft64(v1, -int(v2)))
}

func (vm *VM) i64Eq() {
	vm.pushBool(vm.popUint64() == vm.popUint64())
}

func (vm *VM) i64Eqz() {
	vm.pushBool(vm.popUint64() == 0)
}

func (vm *VM) i64Ne() {
	vm.pushBool(vm.popUint64() != vm.popUint64())
}

func (vm *VM) i64LtS() {
	v2 := vm.popInt64()
	v1 := vm.popInt64()
	vm.pushBool(v1 < v2)
}

func (vm *VM) i64LtU() {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushBool(v1 < v2)
}

func (vm *VM) i64GtS() {
	v2 := vm.popInt64()
	v1 := vm.popInt64()
	vm.pushBool(v1 > v2)
}

func (vm *VM) i64GtU() {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushBool(v1 > v2)
}

func (vm *VM) i64LeU() {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushBool(v1 <= v2)
}

func (vm *VM) i64LeS() {
	v2 := vm.popInt64()
	v1 := vm.popInt64()
	vm.pushBool(v1 <= v2)
}

func (vm *VM) i64GeS() {
	v2 := vm.popInt64()
	v1 := vm.popInt64()
	vm.pushBool(v1 >= v2)
}

func (vm *VM) i64GeU() {
	v2 := vm.popUint64()
	v1 := vm.popUint64()
	vm.pushBool(v1 >= v2)
}

// float32 operators

func (vm *VM) f32Abs() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat32()
	val := float32(math.Abs(float64(v1)))
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x8B, "f32 Absolute", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Neg() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat32()
	val := -v1
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x8C, "f32 Negative", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Ceil() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat32()
	val := float32(math.Ceil(float64(v1)))
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x8D, "f32 Ceiling", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Floor() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat32()
	val := float32(math.Floor(float64(v1)))
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x8E, "f32 Floor", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Trunc() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat32()
	val := float32(math.Trunc(float64(v1)))
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x8F, "f32 Trunc", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Nearest() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	f := vm.popFloat32()
	val := float32(int32(f + float32(math.Copysign(0.5, float64(f)))))
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x90, "f32 Nearest", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, f, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Sqrt() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v1 := vm.popFloat32()
	val := float32(math.Sqrt(float64(v1)))
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x91, "f32 Square root", []string{"program_counter", "base_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Add() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	val := v1 + v2
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x92, "f32 Add", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Sub() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	val := v1 - v2
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x93, "f32 Sub", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Mul() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	val := v1 * v2
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x94, "f32 Multiply", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Div() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	val := v1 / v2
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x95, "f32 Divide", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Min() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	val := float32(math.Min(float64(v1), float64(v2)))
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x96, "f32 Min", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Max() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	val := float32(math.Max(float64(v1), float64(v2)))
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x97, "f32 Max", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Copysign() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	val := float32(math.Copysign(float64(v1), float64(v2)))
	vm.pushFloat32(val)

	// Log this operation
	opLog(vm, 0x98, "f32 Copy sign", []string{"program_counter", "base_value", "modifier_value", "result_value", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, val, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Eq() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	cond := v1 == v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x5B, "f32 Equal", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Ne() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	cond := v1 != v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x5C, "f32 Not equal", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Lt() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	cond := v1 < v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x5D, "f32 Less than", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Gt() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	cond := v1 > v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x5C, "f32 Greater than", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Le() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	cond := v1 <= v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x5F, "f32 Less than or equal", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

func (vm *VM) f32Ge() {
	stackStart := vm.ctx.stack

	// The operation we're logging
	v2 := vm.popFloat32()
	v1 := vm.popFloat32()
	cond := v1 >= v2
	vm.pushBool(cond)

	// Log this operation
	opLog(vm, 0x60, "f32 Greater than or equal", []string{"program_counter", "base_value", "modifier_value", "condition_met", "stack_start", "stack_finish"},
		[]interface{}{vm.ctx.pc, v1, v2, cond, stackStart, vm.ctx.stack})
}

// float64 operators

func (vm *VM) f64Abs() {
	vm.pushFloat64(math.Abs(vm.popFloat64()))
}

func (vm *VM) f64Neg() {
	vm.pushFloat64(-vm.popFloat64())
}

func (vm *VM) f64Ceil() {
	vm.pushFloat64(math.Ceil(vm.popFloat64()))
}

func (vm *VM) f64Floor() {
	vm.pushFloat64(math.Floor(vm.popFloat64()))
}

func (vm *VM) f64Trunc() {
	vm.pushFloat64(math.Trunc(vm.popFloat64()))
}

func (vm *VM) f64Nearest() {
	f := vm.popFloat64()
	vm.pushFloat64(float64(int64(f + math.Copysign(0.5, f))))
}

func (vm *VM) f64Sqrt() {
	vm.pushFloat64(math.Sqrt(vm.popFloat64()))
}

func (vm *VM) f64Add() {
	vm.pushFloat64(vm.popFloat64() + vm.popFloat64())
}

func (vm *VM) f64Sub() {
	v2 := vm.popFloat64()
	v1 := vm.popFloat64()
	vm.pushFloat64(v1 - v2)
}

func (vm *VM) f64Mul() {
	vm.pushFloat64(vm.popFloat64() * vm.popFloat64())
}

func (vm *VM) f64Div() {
	v2 := vm.popFloat64()
	v1 := vm.popFloat64()
	vm.pushFloat64(v1 / v2)
}

func (vm *VM) f64Min() {
	vm.pushFloat64(math.Min(vm.popFloat64(), vm.popFloat64()))
}

func (vm *VM) f64Max() {
	vm.pushFloat64(math.Max(vm.popFloat64(), vm.popFloat64()))
}

func (vm *VM) f64Copysign() {
	vm.pushFloat64(math.Copysign(vm.popFloat64(), vm.popFloat64()))
}

func (vm *VM) f64Eq() {
	vm.pushBool(vm.popFloat64() == vm.popFloat64())
}

func (vm *VM) f64Ne() {
	vm.pushBool(vm.popFloat64() != vm.popFloat64())
}

func (vm *VM) f64Lt() {
	v2 := vm.popFloat64()
	v1 := vm.popFloat64()
	vm.pushBool(v1 < v2)
}

func (vm *VM) f64Gt() {
	v2 := vm.popFloat64()
	v1 := vm.popFloat64()
	vm.pushBool(v1 > v2)
}

func (vm *VM) f64Le() {
	v2 := vm.popFloat64()
	v1 := vm.popFloat64()
	vm.pushBool(v1 <= v2)
}

func (vm *VM) f64Ge() {
	v2 := vm.popFloat64()
	v1 := vm.popFloat64()
	vm.pushBool(v1 >= v2)
}
