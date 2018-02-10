# JVM中的class字节码解析

## 1. class文件组成

class文件由以下部分组成(按顺序)

1. magic number: 0xCAFEBABE. 标志着文件类型为java的class文件
2. minorVersion: 当前class文件支持的副版本号，从java1.2以后都是0
3. majorVersion: 当前class文件支持的主版本号
4. constantpool: 常量池
5. accessFlags: 类的访问标志,标志着class文件是类还是接口，public还是private
6. thisClass: 当前类全名
8. superClass: 父类全名
9. interfaces: 类实现的所有接口名
10. fields: 类的字段表
11. methods: 类的方法签名表
12. attributes: 类的属性表,代码编译成的字节码就在这里

java虚拟机规定，如果class文件不符合规范，就应该抛出`java.lang.ClassFormatError`

```golang
type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	costantpool  ConstantPool //常量池
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}
```

## 2. 常量池ConstantPool

常量池占据了class文件很大一块数据，存放着各式各样的常量信息，包括数字和字符串常量，类和接口名称，字段和方法名等
java虚拟机规定的常量种类是14种

```golang
const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)
```

需要注意的是`CONSTANT_Utf8`,这种常量存放的是MUTF-8编码的字符串
和UTF-8的差别有两点:

1. null字符(U+0000)会编码成2个字节(byte, u2): `0xC0,0x80`
2. 补充字符(代码点大于U+FFFF的Unicode字符)是按照UTF17拆分为代理对分别编码的

大部分的字符串常量都不直接存值，而是存指向utf8常量的索引

`NameAndType`类型是表示字段或者方法的名称和描述符，描述符的语法规定是:

- 类型描述符
  * 基本类型byte,short,char,int,long,float,double的描述符是单个字母,B,S,C,I,J,F,D
  * 引用类型的描述符是`L+类的完全限定符;`,如`Ljava/lang/Object;`
  * 数组类型的描述符是`[+数组元素类型描述符`,如`[[java/lang/Object`表示Object类型的二维数组
- 字段描述符: 字段类型描述符
- 方法描述符: 分号分隔的参数类型描述符+返回值类型描述符，void返回值由单个字母V表示

## 3. 属性表

属性表存放着比较杂的东西，如方法的字节码.
属性是可以扩展的

按照用途，23种预定义属性可以分为3种:

1. 实现虚拟机所必须的:`ConstantValue`,`Code`,`Exceptions`,`StackMapTable`,`BootstrapMethods`
2. 实现Java类库必须的
3. 提供给工具使用


![](http://7xox4k.com1.z0.glb.clouddn.com/2018/02/10/20180210200638.png)
![](http://7xox4k.com1.z0.glb.clouddn.com/2018/02/10/20180210200747.png)
