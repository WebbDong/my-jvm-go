package jcmd

import (
	"flag"
	"fmt"
	"os"
)

// java 命令行工具
type Java struct {
	HelpFlag    bool
	VersionFlag bool
	CpOption    string
	Class       string
	Args        []string
}

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

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
