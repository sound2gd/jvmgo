package extended

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
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
