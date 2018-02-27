package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
)

type DNEG struct {
	base.NoOperandsInstruction
}

func (self *DNEG) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v := stack.PopDouble()
	stack.PushDouble(-v)
}

type FNEG struct {
	base.NoOperandsInstruction
}

func (self *FNEG) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v := stack.PopFloat()
	stack.PushFloat(-v)
}

type INEG struct {
	base.NoOperandsInstruction
}

func (self *INEG) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v := stack.PopInt()
	stack.PushInt(-v)
}

type LNEG struct {
	base.NoOperandsInstruction
}

func (self *LNEG) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v := stack.PopLong()
	stack.PushLong(-v)
}
