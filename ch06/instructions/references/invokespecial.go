package references

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (self *INVOKE_SPECIAL) Execute(frame *rtdata.Frame) {
	frame.OperandStack().PopRef()
}
