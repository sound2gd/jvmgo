package extended

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
)

// 判断引用为null
type IFNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtdata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFNONNULL) Execute(frame *rtdata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}
