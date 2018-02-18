package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

/*
从操作数栈读取变量，然后放入局部变量表
*/
type ISTORE struct {
	base.Index8Instruction
}

func (self *ISTORE) Execute(frame *rtdata.Frame) {
	_istore(frame, self.Index)
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_0) Execute(frame *rtdata.Frame) {
	_istore(frame, 0)
}

type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_1) Execute(frame *rtdata.Frame) {
	_istore(frame, 1)
}

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_2) Execute(frame *rtdata.Frame) {
	_istore(frame, 2)
}

type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_3) Execute(frame *rtdata.Frame) {
	_istore(frame, 3)
}

func _istore(frame *rtdata.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}
