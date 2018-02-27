package stack

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
)

type POP struct {
	base.NoOperandsInstruction
}

func (self POP) Execute(frame *rtdata.Frame){
	frame.OperandStack().PopSlot()
}


// 对于double和long,占用2个操作数栈slot，需要使用POP2指令
type POP2 struct{
	base.NoOperandsInstruction
}

func (self POP2) Execute(frame *rtdata.Frame){
	operandStack := frame.OperandStack()
	operandStack.PopSlot()
	operandStack.PopSlot()
}
