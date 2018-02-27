package heap

import "jvmgo/ch06/classfile"

type FieldRef struct {
	MemberRef
	// 缓存解析后的字段指针
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldRefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberRefInfo)
	return ref
}

func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}
func (self *FieldRef) resolveFieldRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	field := lookupField(c, self.name, self.descripter)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field

}

// 从类中寻找字段
func lookupField(class *Class, name, descriptor string) *Field {
	// 从本类查找
	for _, field := range class.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	// 从接口里查找
	for _, iface := range class.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	// 从父类里查找
	if class.superClass != nil {
		return lookupField(class.superClass, name, descriptor)
	}

	return nil
}
