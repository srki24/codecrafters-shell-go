package main

import (
	"fmt"
	"os"
	"path/filepath"
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

	paths := filepath.SplitList(os.Getenv("PATH"))

	for _, path := range paths {
		fp := filepath.Join(path, cmd)

		if _, err := os.Stat(fp); err == nil {
			fmt.Printf("%s is %s\n", cmd, fp)
			return
		}
	}

	fmt.Printf("%s: not found\n", cmd)
}
