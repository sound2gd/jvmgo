# JVM如何寻找class文件

JVM规范并没有规定从哪里去寻找class文件, 不同的虚拟机有不同的实现
Oracle的虚拟机实现是根据classpath来搜索类,按照搜索的先后顺序,可以分为:

1. BootstrapClasspath: 启动类路径,默认在`$JAVA_HOME/jre/lib`，java的标准库(大部分在rt.jar)位于该路径
2. ExtensionClasspath: 扩展类加载路径,默认在`$JAVA_Home/jre/lib/ext`
3. User Classpath: 用户类路径, 自己实现的类和第三方类库都在这,默认值是`.`


在执行的时候，通过`-classpath/-cp`可以指定classpath,既可以是目录，也可以是jar文件,java6开始支持通配符

本章的实现使用的是组合模式,按照上面的思想实现，读取出某类的class字节码，最后效果图:
![](http://7xox4k.com1.z0.glb.clouddn.com/2018/02/08/20180208162717.png)
