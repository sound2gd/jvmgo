package conversions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

type L2F struct {
	base.NoOperandsInstruction
}

func (self *L2F) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopLong()
	d := float32(f)
	stack.PushFloat(d)
}

type L2I struct {
	base.NoOperandsInstruction
}

func (self *L2I) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopLong()
	d := int32(f)
	stack.PushInt(d)
}

type L2D struct {
	base.NoOperandsInstruction
}

func (self *L2D) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopLong()
	d := int64(f)
	stack.PushLong(d)
}
