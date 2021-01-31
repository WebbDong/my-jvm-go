package main

import (
	"classloader"
	"fmt"
	"jcmd"
	"os"
)

const JavaHomeEnvName = "JAVA_HOME"

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
	loadClass(javaCmd)
}

func loadClass(javaCmd *jcmd.Java) {
	var javaHome string
	if javaCmd.Xbootcp != "" {
		javaHome = javaCmd.Xbootcp
	} else {
		javaHome = os.Getenv(JavaHomeEnvName)
	}

	rtJarPath := javaHome + "/jre/lib/rt.jar"
	jarClassLoader := classloader.JarClassLoader{AbsJarPath: rtJarPath}
	data, err := jarClassLoader.LoadClass("java/lang/Object")
	if err != nil {
		panic(err)
	}
	fmt.Printf("class data: %v\n", data)
}
