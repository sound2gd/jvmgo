package comparisons

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
)

type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *rtdata.Frame) {
	if _acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ACMPNE) Execute(frame *rtdata.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

func _acmp(frame *rtdata.Frame) bool {
	stack := frame.OperandStack()
	v2 := stack.PopRef()
	v1 := stack.PopRef()
	return v1 == v2
}
