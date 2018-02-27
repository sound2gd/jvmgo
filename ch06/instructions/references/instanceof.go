package references

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
	"jvmgo/ch06/rtdata/heap"
)

// 判断对象是否是某个类的实例(或者对象的类是否实现了某个接口),并且把结果推到操作数栈
// 第一个操作数是方法的字节码里获取的uint16索引
// 第二个操作数是对象引用,从操作数栈弹出
type INSTANCE_OF struct{ base.Index16Instruction }

func (self *INSTANCE_OF) Execute(frame *rtdata.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolveClass()

	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
