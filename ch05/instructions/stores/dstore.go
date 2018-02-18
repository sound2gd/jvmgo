package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

/*
从操作数栈读取变量，然后放入局部变量表
*/
type DSTORE struct {
	base.Index8Instruction
}

func (self *DSTORE) Execute(frame *rtdata.Frame) {
	_dstore(frame, self.Index)
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_0) Execute(frame *rtdata.Frame) {
	_dstore(frame, 0)
}

type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_1) Execute(frame *rtdata.Frame) {
	_dstore(frame, 1)
}

type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_2) Execute(frame *rtdata.Frame) {
	_dstore(frame, 2)
}

type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_3) Execute(frame *rtdata.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *rtdata.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
