package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	// java 命令运行的参数 -help -? -cp -version -classpath
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help string")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help string")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...] \n", os.Args[0])
}
