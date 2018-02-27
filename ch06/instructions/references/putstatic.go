package references

import (
	"jvmgo/ch06/instructions/base"
	"jvmgo/ch06/rtdata"
	"jvmgo/ch06/rtdata/heap"
)

// 设置class的静态属性
type PUT_STATIC struct{ base.Index16Instruction }

func (self *PUT_STATIC) Execute(frame *rtdata.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	if field.IsFinal() {
		// 如果是final类型，那么只有在类初始化的时候给它赋值
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descripter()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	}
}
