package rtdata

import "jvmgo/ch06/rtdata/heap"

/*
 Frame表示栈帧
 这里栈帧的实现是链表
 栈帧里有局部变量表，操作数栈
*/
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC       int // 下一个调用地址
}

func NewFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

// getters
func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPc(nextPC int) {
	self.nextPC = nextPC
}
