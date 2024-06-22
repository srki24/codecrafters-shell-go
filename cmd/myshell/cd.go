package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func go_cd(argv ...string) {

	path := argv[0]
	if path[0] == '~' {
		homedir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		path = strings.Replace(path, "~", homedir, 1)
	}

	err := os.Chdir(path)

	if err != nil {
		fmt.Printf("cd: %s: No such file or directory\n", path)
	}

}
