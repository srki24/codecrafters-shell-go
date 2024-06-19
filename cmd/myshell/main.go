package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		text = strings.TrimSpace(text)
		fmt.Printf("%s: command not found\n", text)

	}
}
