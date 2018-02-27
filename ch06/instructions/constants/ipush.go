package constants

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
)

// 从操作数获取一个byte，转换成int,然后推入操作数栈顶
type BIPUSH struct {
	val int8
}

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtdata.Frame){
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

// 读取一个short，转换成int,然后推入操作数栈顶
type SIPUSH struct {
	val int16
}

func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtdata.Frame){
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
