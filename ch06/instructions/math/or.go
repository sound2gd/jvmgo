package math

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
)

// int or boolean or
type IOR struct {
	base.NoOperandsInstruction
}

func (self *IOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	r := v1 | v2
	stack.PushInt(r)
}

type LOR struct {
	base.NoOperandsInstruction
}

func (self *LOR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	r := v1 | v2
	stack.PushLong(r)
}
