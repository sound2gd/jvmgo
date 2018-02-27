package references

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
	"jvmgo/ch06/rtdata/heap"
)

// 检查类型转换，不改变操作数栈
type CHECK_CAST struct{ base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)

	if ref == nil {
		// null可以cast成任意类型
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	if !ref.IsInstanceOf(classRef.ResolveClass()) {
		panic("java.lang.ClassCastException")
	}
}
