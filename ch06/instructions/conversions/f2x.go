package conversions

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
)

type F2D struct {
	base.NoOperandsInstruction
}

func (self *F2D) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

type F2I struct {
	base.NoOperandsInstruction
}

func (self *F2I) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := int32(f)
	stack.PushInt(d)
}

type F2L struct {
	base.NoOperandsInstruction
}

func (self *F2L) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := int64(f)
	stack.PushLong(d)
}
