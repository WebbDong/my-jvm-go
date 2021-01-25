package classloader

import (
	"errors"
	"strings"
)

// 类加载器切片数组，多个类加载器的组合
type CompositeClassLoader []ClassLoader

func CreateCompositeClassLoader(pathList string) CompositeClassLoader {
	compositeClassLoader := CompositeClassLoader{}
	paths := strings.Split(pathList, pathListSeparator)
	for _, path := range paths {
		compositeClassLoader = append(compositeClassLoader, NewClassLoader(path))
	}
	return compositeClassLoader
}

func (s CompositeClassLoader) loadClass(classname string) ([]byte, error) {
	for _, cl := range s {
		data, err := cl.loadClass(classname)
		if err == nil {
			return data, nil
		}
	}
	return nil, errors.New("class not found: " + classname)
}
