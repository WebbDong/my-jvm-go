package classloader

import (
	"io/ioutil"
	"path/filepath"
)

// 读取加载目录中的 class 文件的 ClassLoader
type DirClassLoader struct {
	AbsDirPath string
}

func CreateDirClassLoader(path string) *DirClassLoader {
	// 将路径转换成绝对路径，如果 path 是相对路径，生成的绝对路径就是执行的程序文件所在的目录为根目录
	absDirPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirClassLoader{AbsDirPath: absDirPath}
}

func (d *DirClassLoader) loadClass(classname string) ([]byte, error) {
	absFilePath := filepath.Join(d.AbsDirPath, classname)
	data, err := ioutil.ReadFile(absFilePath)
	return data, err
}
