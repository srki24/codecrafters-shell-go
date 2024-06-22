package main

import (
	"fmt"
	"log"
	"os"
)

func go_pwd(argv ...string) {

	cwd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cwd)
}
