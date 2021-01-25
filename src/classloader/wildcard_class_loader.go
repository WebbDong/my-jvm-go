package classloader

import (
	"os"
	"path/filepath"
	"strings"
)

func CreateWildcardClassLoader(path string) CompositeClassLoader {
	// 去除末尾的通配符 *
	baseDir := path[:len(path)-1]
	compositeClassLoader := CompositeClassLoader{}
	filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") ||
			strings.HasSuffix(path, ".JAR") {
			jarClassLoader := CreateJarClassLoader(path)
			compositeClassLoader = append(compositeClassLoader, jarClassLoader)
		}
		return nil
	})
	return nil
}
