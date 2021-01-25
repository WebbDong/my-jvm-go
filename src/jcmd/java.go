package jcmd

import (
	"flag"
	"fmt"
	"os"
)

// java 命令行工具
type Java struct {
	// -help
	HelpFlag bool
	// -version
	VersionFlag bool
	// -classpath or -cp
	CpOption       string
	Xbootclasspath string
	// 完整全类名
	Class string
	// 命令行参数
	Args []string
}

// 解析 flag 命令参数
func ParseJavaCmd() *Java {
	java := &Java{}
	flag.Usage = PrintUsage
	flag.BoolVar(&java.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&java.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&java.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&java.CpOption, "classpath", "", "classpath")
	flag.StringVar(&java.CpOption, "cp", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		java.Class = args[0]
		java.Args = args[1:]
	}
	return java
}

// 打印命令行使用说明
func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
