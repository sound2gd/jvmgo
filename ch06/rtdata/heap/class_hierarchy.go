package heap

func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self
	if s == t {
		return true
	}

	// T t = (T) s
	if !t.IsInterface() {
		// 如果s是t的子类,那么s可以赋值给t(向上转型)
		return s.isSubClassOf(t)
	} else {
		// 如果t是接口，那么需要s实现了t
		return s.isImplements(t)
	}
}
func (self *Class) isSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) isImplements(other *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == other || i.isSubInterfaceOf(other) {
				return true
			}
		}
	}
	return false
}
func (self *Class) isSubInterfaceOf(other *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == other ||
			superInterface.isSubInterfaceOf(other) {
			return true
		}
	}
	return false
}
