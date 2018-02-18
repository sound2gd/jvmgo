package comparisons

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ICMPEQ) Execute(frame *rtdata.Frame) {
	if v1, v2 := _icmpPop(frame); v1 == v2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPNE struct {
	base.BranchInstruction
}

func (self *IF_ICMPNE) Execute(frame *rtdata.Frame) {
	if v1, v2 := _icmpPop(frame); v1 != v2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPLT struct {
	base.BranchInstruction
}

func (self *IF_ICMPLT) Execute(frame *rtdata.Frame) {
	if v1, v2 := _icmpPop(frame); v1 < v2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPLE struct {
	base.BranchInstruction
}

func (self *IF_ICMPLE) Execute(frame *rtdata.Frame) {
	if v1, v2 := _icmpPop(frame); v1 <= v2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPGT struct {
	base.BranchInstruction
}

func (self *IF_ICMPGT) Execute(frame *rtdata.Frame) {
	if v1, v2 := _icmpPop(frame); v1 > v2 {
		base.Branch(frame, self.Offset)
	}
}

type IF_ICMPGE struct {
	base.BranchInstruction
}

func (self *IF_ICMPGE) Execute(frame *rtdata.Frame) {
	if v1, v2 := _icmpPop(frame); v1 >= v2 {
		base.Branch(frame, self.Offset)
	}
}

func _icmpPop(frame *rtdata.Frame) (v1, v2 int32) {
	stack := frame.OperandStack()
	v2 = stack.PopInt()
	v1 = stack.PopInt()
	return
}
