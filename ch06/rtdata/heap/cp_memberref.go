package heap

import "jvmgo/ch06/classfile"

type MemberRef struct {
	SymRef
	name string
	// java中不允许同一个类中定义名称相同，但是类型不同的两个字段，从虚拟机实现来看是可以做到的
	descripter string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberRefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descripter = refInfo.NameAndDescriptor()
}

func (self *MemberRef) Name() string {
	return self.name
}

func (self *MemberRef) Descriptor() string {
	return self.descripter
}
