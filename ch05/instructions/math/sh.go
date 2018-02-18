package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

// 移位操作
type ISHL struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	// 0x1f就是31，int类型最长32位，取前5位bit表示移位操作就够了
	s := uint32(v1) & 0x1f
	result := v2 << s
	stack.PushInt(result)
}

type ISHR struct{ base.NoOperandsInstruction }

func (self *ISHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()

	s := uint32(v1) & 0x1f
	result := v2 >> s
	stack.PushInt(result)
}

type IUSHR struct {
	base.NoOperandsInstruction
}

func (self *IUSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	s := uint32(v1) & 0x1f
	// golang里面没有无符号移位，只好转换成uint移位完成以后再转成int
	result := uint32(v2) >> s
	stack.PushInt(int32(result))
}

type IUSHL struct {
	base.NoOperandsInstruction
}

func (self *IUSHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	s := uint32(v1) & 0x1f
	result := uint32(v2) << s
	stack.PushInt(int32(result))
}

type LSHL struct {
	base.NoOperandsInstruction
}

func (self *LSHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	s := uint64(v1) & 0x3f
	result := v2 << s
	stack.PushLong(result)
}

type LSHR struct {
	base.NoOperandsInstruction
}

func (self *LSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	s := uint64(v1) & 0x3f
	result := v2 >> s
	stack.PushLong(result)
}

type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *LUSHR) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	s := uint64(v1) & 0x3f
	result := uint64(v2) >> s
	stack.PushLong(int64(result))
}

type LUSHL struct {
	base.NoOperandsInstruction
}

func (self *LUSHL) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	s := uint64(v1) & 0x3f
	result := uint64(v2) << s
	stack.PushLong(int64(result))
}
