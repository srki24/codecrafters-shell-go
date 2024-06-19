package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	var text string

	_, err := fmt.Scanln(&text)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s: command not found", text)
	// Wait for user input
	bufio.NewReader(os.Stdin).ReadString('\n')
}
