package heap

// 符号引用，通用数据结构
type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

func (self *SymRef) ResolveClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	// 如果类符号已经解析，直接返回类指针
	return self.class
}

func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
