package main

import (
	"fmt"
	"strings"
)

func go_echo(argv ...string) {
	if len(argv) == 0 {
		return
	}

	fmt.Println(strings.Join(argv, " "))
}
