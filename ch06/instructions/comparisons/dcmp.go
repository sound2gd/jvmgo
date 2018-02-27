package comparisons

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
)

type DCMPG struct {
	base.NoOperandsInstruction
}

func (self *DCMPG) Execute(frame *rtdata.Frame) {
	_dcmp(frame, true)
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (self *DCMPL) Execute(frame *rtdata.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtdata.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else {
		// 浮点数比较会发生NaN,无法比较
		if gFlag {
			stack.PushInt(1)
		} else {
			stack.PushInt(-1)
		}
	}
}
