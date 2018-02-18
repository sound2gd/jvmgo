package extended

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

// 和Goto一样，只是这次读4字节
type GOTO_W struct {
	offset int
}

func (self *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtdata.Frame) {
	base.Branch(frame, self.offset)
}
