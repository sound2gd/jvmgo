package conversions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

type D2F struct {
	base.NoOperandsInstruction
}

func (self *D2F) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopDouble()
	d := float32(f)
	stack.PushFloat(d)
}

type D2I struct {
	base.NoOperandsInstruction
}

func (self *D2I) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopDouble()
	d := int32(f)
	stack.PushInt(d)
}

type D2L struct {
	base.NoOperandsInstruction
}

func (self *D2L) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopDouble()
	d := int64(f)
	stack.PushLong(d)
}
