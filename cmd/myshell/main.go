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
	_, err := fmt.Scan(&text)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s: command not found\n", text)
	// Wait for user input
	bufio.NewReader(os.Stdin).ReadString('\n')
}
