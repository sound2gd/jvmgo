package references

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
	"jvmgo/ch06/rtdata/heap"
)

// 给实例变量赋值，需要3个操作数，前俩操作数是常量池索引和变量值
// 第三个操作数是对象引用
type PUT_FIELD struct{ base.Index16Instruction }

func (self *PUT_FIELD) Execute(frame *rtdata.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	if field.IsFinal() {
		// 如果是final类型，那么只有在类初始化的时候给它赋值
		if currentClass != class || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descripter()
	slotId := field.SlotId()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	}
}
