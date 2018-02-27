package rtdata

import "jvmgo/ch06/rtdata/heap"

/*
 Thread
   pc
   Stack
     Frame
      LocalVars
      OperandStack
*/
type Thread struct {
	// 当前执行的指令位置
	pc int
	// 栈
	stack *Stack
	// TODO 后面完善
}

func NewThread() *Thread {
	return &Thread{
		// 默认栈大小1024
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return NewFrame(self, method)
}
