package heap

import (
	"fmt"
	"jvmgo/ch06/classfile"
)

// 运行时常量池主要存放两类信息：字面量(literal)和符号引用(Symbolic reference)
// 字面量包括整数，浮点数和字符串字面量
// 符号引用包括类符号引用，字段符号引用和接口方法符号引用
type Constant interface{}
type ConstantPool struct {
	class  *Class
	consts []Constant
}

func newConstantPool(class *Class, cfCp classfile.ConstantPool) *ConstantPool {
	cpCount := len(cfCp)
	consts := make([]Constant, cpCount)
	rtCp := &ConstantPool{class, consts}

	for i := 1; i < cpCount; i++ {
		cpInfo := cfCp[i]
		switch cpInfo.(type) {
		// int和float型常亮，直接取出常量值放入
		case *classfile.ConstantIntegerInfo:
			intInfo := cpInfo.(*classfile.ConstantIntegerInfo)
			consts[i] = intInfo.Value()
		case *classfile.ConstantFloatInfo:
			floatInfo := cpInfo.(*classfile.ConstantFloatInfo)
			consts[i] = floatInfo.Value()
		// Long和Double要占俩字节
		case *classfile.ConstantLongInfo:
			longInfo := cpInfo.(*classfile.ConstantLongInfo)
			consts[i] = longInfo.Value()
			i++
		case *classfile.ConstantDoubleInfo:
			doubleInfo := cpInfo.(*classfile.ConstantDoubleInfo)
			consts[i] = doubleInfo.Value()
			i++
		case *classfile.ConstantStringInfo:
			strInfo := cpInfo.(*classfile.ConstantStringInfo)
			consts[i] = strInfo.String()

		// 下面都是符号引用
		case *classfile.ConstantClassInfo:
			classInfo := cpInfo.(*classfile.ConstantClassInfo)
			consts[i] = newClassRef(rtCp, classInfo)
		case *classfile.ConstantFieldRefInfo:
			fieldrefInfo := cpInfo.(*classfile.ConstantFieldRefInfo)
			consts[i] = newFieldRef(rtCp, fieldrefInfo)
		case *classfile.ConstantMethodRefInfo:
			methodrefInfo := cpInfo.(*classfile.ConstantMethodRefInfo)
			consts[i] = newMethodRef(rtCp, methodrefInfo)
		case *classfile.ConstantInterfaceMethodRefInfo:
			inrefInfo := cpInfo.(*classfile.ConstantInterfaceMethodRefInfo)
			consts[i] = newInterfaceMethodRef(rtCp, inrefInfo)
		default:
			// TODO
		}
	}
	return rtCp
}

func (self *ConstantPool) GetConstant(index uint) Constant {
	if c := self.consts[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("No constants at index %d", index))
}
