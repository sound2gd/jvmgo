package constants

import "jvmgo/ch06/instructions/base"
import "jvmgo/ch06/rtdata"

// 空指令
type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtdata.Frame) {
	// 空指令不用做啥
}
