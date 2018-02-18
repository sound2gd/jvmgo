package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

// 交换栈顶的俩元素
type SWAP struct {
	base.NoOperandsInstruction
}

func (self *SWAP) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	x1 := stack.PopSlot()
	x2 := stack.PopSlot()

	stack.PushSlot(x1)
	stack.PushSlot(x2)
}
