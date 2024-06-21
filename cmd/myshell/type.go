package main

import (
	"fmt"
)

func go_type(argv []string) {
	cmd := ""
	if len(argv) > 1 {
		cmd = argv[1]
	}

	if hasFunction(cmd) {
		fmt.Printf("%s is a shell builtin\n", cmd)
		return
	}

	fmt.Printf("%s: not found\n", cmd)
}
