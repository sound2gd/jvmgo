package control

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

// 纯跳转指令
type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtdata.Frame) {
	// 无条件跳转
	base.Branch(frame, self.Offset)
}


