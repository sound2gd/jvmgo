package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtdata"
)

// DUP用于将操作数栈顶的变量复制一份
type DUP struct{ base.NoOperandsInstruction }

func (self DUP) Execute(frame *rtdata.Frame) {
	opStack := frame.OperandStack()
	slot := opStack.PopSlot()
	opStack.PushSlot(slot)
	opStack.PushSlot(slot)
}

type DUP_X1 struct{ base.NoOperandsInstruction }

func (self DUP_X1) Execute(frame *rtdata.Frame) {
	operandStack := frame.OperandStack()
	x1 := operandStack.PopSlot()
	x2 := operandStack.PopSlot()
	operandStack.PushSlot(x1)
	operandStack.PushSlot(x2)
	operandStack.PushSlot(x1)
}

type DUP2 struct{ base.NoOperandsInstruction }


func (self DUP2) Execute(frame *rtdata.Frame) {

}

type DUP_X2 struct{ base.NoOperandsInstruction }

func (self DUP_X2) Execute(frame *rtdata.Frame) {
	operandStack := frame.OperandStack()
	x1 := operandStack.PopSlot()
	x2 := operandStack.PopSlot()
	x3 := operandStack.PopSlot()
	operandStack.PushSlot(x1)
	operandStack.PushSlot(x3)
	operandStack.PushSlot(x2)
	operandStack.PushSlot(x1)
}

type DUP2_X1 struct {
	base.NoOperandsInstruction
}

func (self DUP2_X1) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	x1 := stack.PopSlot()
	x2 := stack.PopSlot()
	x3 := stack.PopSlot()
	stack.PushSlot(x2)
	stack.PushSlot(x1)
	stack.PushSlot(x3)
	stack.PushSlot(x2)
	stack.PushSlot(x1)
}

type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (self DUP2_X2) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	x1 := stack.PopSlot()
	x2 := stack.PopSlot()
	x3 := stack.PopSlot()
	x4 := stack.PopSlot()
	stack.PushSlot(x2)
	stack.PushSlot(x1)
	stack.PushSlot(x4)
	stack.PushSlot(x3)
	stack.PushSlot(x2)
	stack.PushSlot(x1)
}
