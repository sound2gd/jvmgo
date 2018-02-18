package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

type IAND struct {
	base.NoOperandsInstruction
}

func (self *IAND) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	stack.PushInt(v1 & v2)
}
type LAND struct {
	base.NoOperandsInstruction
}

func (self *LAND) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	stack.PushLong(v1 & v2)
}
