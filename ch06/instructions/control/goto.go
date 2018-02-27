package control

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
)

// 纯跳转指令
type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtdata.Frame) {
	// 无条件跳转
	base.Branch(frame, self.Offset)
}


