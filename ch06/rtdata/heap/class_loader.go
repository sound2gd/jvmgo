package heap

import (
	"fmt"
	"jvmgo/ch06/classfile"
	"jvmgo/ch06/classpath"
)

/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/
type ClassLoader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class // 缓存已经加载的类
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

func (self *ClassLoader) LoadClass(name string) *Class {
	// 检查类是否已经被加载，如果已经被加载，直接返回类数据
	if class, ok := self.classMap[name]; ok {
		return class
	}
	// 数组类和普通类有很大的不同，数组类的数据由虚拟机直接生成
	return self.loadNonArrayClass(name)
}

// 类的加载大致分成3个步骤，首先找到class文件并且把数据读取到内存,然后解析class文件,生成虚拟机可以使用的类数据
// 最后进行链接
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newCLass(cf)
}

func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

// 类的链接分为验证和准备俩阶段
func link(class *Class) {
	verify(class)
	prepare(class)
}

func verify(class *Class) {
	// TODO
}

func prepare(class *Class) {
	// 给类变量分配空间并且给予初始值
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

/*
  计算实例字段的个数，给实例字段编号
  实例字段需要递归的查找父类
*/
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}

	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}

	class.instanceSlotCount = slotId
}

/*
  计算静态字段的个数,同时给静态字段编号
*/
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			// 给静态字段编号
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	slotId := field.slotId
	cpIndex := field.slotId
	if cpIndex > 0 {
		switch field.Descripter() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J": // Long型
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("TODO")
		}
	}
}
