package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
)

type DMUL struct {
	base.NoOperandsInstruction
}

func (self *DMUL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	r := v1 * v2
	stack.PushDouble(r)
}

type FMUL struct {
	base.NoOperandsInstruction
}

func (self *FMUL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	r := v1 * v2
	stack.PushFloat(r)
}

type IMUL struct {
	base.NoOperandsInstruction
}

func (self *IMUL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	r := v1 * v2
	stack.PushInt(r)
}

type LMUL struct {
	base.NoOperandsInstruction
}

func (self *LMUL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	r := v1 * v2
	stack.PushLong(r)
}
