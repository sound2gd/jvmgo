package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootstrapClasspath Entry
	extClasspath       Entry
	userClasspath      Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootstrapAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseBootstrapAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/* 启动类路径
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootstrapClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/* 扩展类加载路径
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		// 默认用户类加载路径使用的是当前目录
		cpOption = "."
	}

	self.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("can not find folder")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootstrapClasspath.readClass(className); err == nil {
		return data, entry, err
	}

	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}
