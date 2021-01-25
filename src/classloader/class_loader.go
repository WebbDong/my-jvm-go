package classloader

import (
	"os"
	"strings"
)

// 路径分隔符
const pathListSeparator = string(os.PathListSeparator)

// 类加载器接口
type ClassLoader interface {
	// 加载字节码文件
	loadClass(classname string) ([]byte, error)
}

func NewClassLoader(path string) ClassLoader {
	var classLoader ClassLoader
	if strings.Contains(path, pathListSeparator) {
		classLoader = CreateCompositeClassLoader(path)
	} else if strings.HasSuffix(path, "*") {
		classLoader = CreateWildcardClassLoader(path)
	} else if strings.HasSuffix(path, ".jar") ||
		strings.HasSuffix(path, ".JAR") {
		classLoader = CreateJarClassLoader(path)
	} else {
		classLoader = CreateDirClassLoader(path)
	}
	return classLoader
}
