package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

/*
从操作数栈读取变量，然后放入局部变量表
*/
type ASTORE struct {
	base.Index8Instruction
}

func (self *ASTORE) Execute(frame *rtdata.Frame) {
	_astore(frame, self.Index)
}

type ASTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_0) Execute(frame *rtdata.Frame) {
	_astore(frame, 0)
}

type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_1) Execute(frame *rtdata.Frame) {
	_astore(frame, 1)
}

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_2) Execute(frame *rtdata.Frame) {
	_astore(frame, 2)
}

type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_3) Execute(frame *rtdata.Frame) {
	_astore(frame, 3)
}

func _astore(frame *rtdata.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}
