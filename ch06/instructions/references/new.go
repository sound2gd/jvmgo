package references

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
	"jvmgo/ch06/rtdata/heap"
)

// new指令专门用来创建类实例,数组由别的专门的指令创建
type NEW struct {
	// 操作数2个字节，表示类符号引用
	base.Index16Instruction
}

func (self *NEW) Execute(frame *rtdata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolveClass()
	// 接口和抽象类都不能实例化
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
