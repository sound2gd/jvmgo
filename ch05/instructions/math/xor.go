package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

// int or boolean or
type IXOR struct {
	base.NoOperandsInstruction
}

func (self *IXOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	r := v1 ^ v2
	stack.PushInt(r)
}

type LXOR struct {
	base.NoOperandsInstruction
}

func (self *LXOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	r := v1 ^ v2
	stack.PushLong(r)
}
