package heap

import "jvmgo/ch06/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descripter()
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}

func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

func (self *ClassMember) Name() string {
	return self.name
}

func (self *ClassMember) Descripter() string {
	return self.descriptor
}

func (self *ClassMember) Class() *Class {
	return self.class
}

func (self *ClassMember) isAccessibleTo(d *Class) bool {
	// 判断当前类对于d类有没有访问权限
	if self.IsPublic() {
		return true
	}

	// protected 允许同一个包下的访问和子类访问
	c := self.class
	if self.IsProtected() {
		return d == c ||
			d.isSubClassOf(c) ||
			d.getPackageName() == c.getPackageName()
	}

	// default 允许同一个包下的访问
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}

	// private 只有本类自己可以访问
	return d == c
}
