package main

import (
	"fmt"
	"log"
	"os"
)

func exit(argv []string) {
	if len(argv) == 1 {
		os.Exit(0)
	}

	var i int
	_, err := fmt.Sscan(argv[1], &i)

	if err != nil {
		log.Fatal((err))
	}
	os.Exit(i)
}
