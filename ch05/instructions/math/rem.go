package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
	"math"
)

type DREM struct {
	base.NoOperandsInstruction
}

func (self *DREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	x1 := stack.PopDouble()
	x2 := stack.PopDouble()

	result := math.Mod(x2, x1)
	stack.PushDouble(result)
}

type FREM struct {
	base.NoOperandsInstruction
}

func (self *FREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	x1 := stack.PopFloat()
	x2 := stack.PopFloat()

	result := math.Mod(float64(x2), float64(x1))
	stack.PushFloat(float32(result))
}

type IREM struct {
	base.NoOperandsInstruction
}

func (self *IREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	x1 := stack.PopInt()
	x2 := stack.PopInt()

	if x1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := x2 % x1
	stack.PushInt(result)
}

type LREM struct {
	base.NoOperandsInstruction
}

func (self *LREM) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	x1 := stack.PopLong()
	x2 := stack.PopLong()

	if x1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := x2 % x1
	stack.PushLong(result)
}
