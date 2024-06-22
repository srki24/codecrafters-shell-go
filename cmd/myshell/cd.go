package main

import (
	"log"
	"os"
)

func go_cd(argv []string) {

	err := os.Chdir(argv[1])

	if err != nil {
		log.Fatal(err)
	}

}
