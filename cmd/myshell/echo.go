package main

import (
	"fmt"
	"strings"
)

func go_echo(argv []string) {
	if len(argv) == 1 {
		return
	}

	fmt.Println(strings.Join(argv[1:], " "))
}
