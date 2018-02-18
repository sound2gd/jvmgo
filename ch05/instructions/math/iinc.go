package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

type IINC struct {
	// 读取本地变量索引
	Index uint
	// 加上常量
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUInt8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtdata.Frame) {
	vars := frame.LocalVars()
	v := vars.GetInt(self.Index)
	v += self.Const
	vars.SetInt(self.Index, v)
}
