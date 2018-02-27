package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
)

type DADD struct {
	base.NoOperandsInstruction
}

func (self *DADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	r := v1 + v2
	stack.PushDouble(r)
}

type FADD struct {
	base.NoOperandsInstruction
}

func (self *FADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	r := v1 + v2
	stack.PushFloat(r)
}

type IADD struct {
	base.NoOperandsInstruction
}

func (self *IADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	r := v1 + v2
	stack.PushInt(r)
}

type LADD struct {
	base.NoOperandsInstruction
}

func (self *LADD) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	r := v1 + v2
	stack.PushLong(r)
}
