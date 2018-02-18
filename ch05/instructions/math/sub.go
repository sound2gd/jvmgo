package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

type DSUB struct {
	base.NoOperandsInstruction
}

func (self *DSUB) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	r := v1 - v2
	stack.PushDouble(r)
}

type FSUB struct {
	base.NoOperandsInstruction
}

func (self *FSUB) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopFloat()
	v2 := stack.PopFloat()
	r := v1 - v2
	stack.PushFloat(r)
}

type ISUB struct {
	base.NoOperandsInstruction
}

func (self *ISUB) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	r := v1 - v2
	stack.PushInt(r)
}

type LSUB struct {
	base.NoOperandsInstruction
}

func (self *LSUB) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	r := v1 - v2
	stack.PushLong(r)
}
