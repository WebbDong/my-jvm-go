package classloader

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// 读取加载指定 jar 包的 ClassLoader
type JarClassLoader struct {
	AbsJarPath string
}

func CreateJarClassLoader(path string) *JarClassLoader {
	absJarPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &JarClassLoader{AbsJarPath: absJarPath}
}

func (j *JarClassLoader) LoadClass(classname string) ([]byte, error) {
	reader, err := zip.OpenReader(j.AbsJarPath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	classFileName := classname + ".class"
	for _, f := range reader.File {
		if f.Name == classFileName {
			rc, err := f.Open()
			if err != nil {
				rc.Close()
				return nil, err
			}
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				rc.Close()
				return nil, err
			}
			rc.Close()
			return data, err
		}
	}
	return nil, errors.New("class not found: " + classname)
}
