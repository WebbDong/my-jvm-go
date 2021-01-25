package main

import (
	"fmt"
	"jcmd"
)

func main() {
	javaCmd := jcmd.ParseJavaCmd()
	if javaCmd.VersionFlag {
		fmt.Println("version 0.0.1 SNAPSHOT")
	} else if javaCmd.HelpFlag || javaCmd.Class == "" {
		jcmd.PrintUsage()
	} else {
		startJVM(javaCmd)
	}
}

func startJVM(javaCmd *jcmd.Java) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		javaCmd.CpOption, javaCmd.Class, javaCmd.Args)
}
