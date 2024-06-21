package main

import (
	"fmt"
)

func go_type(argv []string) {
	cmd := ""
	if len(argv) > 1 {
		cmd = argv[1]
	}

	if _, ok := getFunction(cmd); ok {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return
	}

	if fp, ok := getExecutablePath(cmd); ok {
		fmt.Printf("%s is %s\n", cmd, fp)
		return
	}

	fmt.Printf("%s: not found\n", cmd)
}
