package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

	paths := strings.Split(os.Getenv("PATH"), ":")

	for _, path := range paths {
		fp := filepath.Join(path, cmd)

		if _, err := os.Stat(fp); err != nil {
			fmt.Printf("%s is %s\n", cmd, path)
			return
		}
	}

	fmt.Printf("%s: not found\n", cmd)
}
