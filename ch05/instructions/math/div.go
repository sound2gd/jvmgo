package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

type DDIV struct {
	base.NoOperandsInstruction
}

func (self *DDIV) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	r := v1 / v2
	stack.PushDouble(r)
}

type FDIV struct {
	base.NoOperandsInstruction
}

func (self *FDIV) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	r := v1 / v2
	stack.PushFloat(r)
}

type IDIV struct {
	base.NoOperandsInstruction
}

func (self *IDIV) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	r := v1 / v2
	stack.PushInt(r)
}

type LDIV struct {
	base.NoOperandsInstruction
}

func (self *LDIV) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	r := v1 / v2
	stack.PushLong(r)
}
