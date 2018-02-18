package conversions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

type I2B struct {
	base.NoOperandsInstruction
}

func (self *I2B) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v := stack.PopInt()
	b := int32(int8(v))
	stack.PushInt(b)
}

type I2C struct {
	base.NoOperandsInstruction
}

func (self *I2C) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v := stack.PopInt()
	b := int32(uint16(v))
	stack.PushInt(b)
}

type I2S struct {
	base.NoOperandsInstruction
}

func (self *I2S) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v := stack.PopInt()
	b := int32(int16(v))
	stack.PushInt(b)
}

type I2L struct {
	base.NoOperandsInstruction
}

func (self *I2L) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v := stack.PopInt()
	b := int64(v)
	stack.PushLong(b)
}

type I2F struct {
	base.NoOperandsInstruction
}

func (self *I2F) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v := stack.PopInt()
	b := float32(v)
	stack.PushFloat(b)
}

type I2D struct {
	base.NoOperandsInstruction
}

func (self *I2D) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v := stack.PopInt()
	b := float64(v)
	stack.PushDouble(b)
}
