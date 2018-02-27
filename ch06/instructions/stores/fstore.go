package stores

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
)

/*
从操作数栈读取变量，然后放入局部变量表
*/
type FSTORE struct {
	base.Index8Instruction
}

func (self *FSTORE) Execute(frame *rtdata.Frame) {
	_fstore(frame, self.Index)
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_0) Execute(frame *rtdata.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_1) Execute(frame *rtdata.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_2) Execute(frame *rtdata.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_3) Execute(frame *rtdata.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *rtdata.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
