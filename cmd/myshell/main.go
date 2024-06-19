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
		cmd, err := bufio.NewReader(os.Stdin).ReadString('\n')

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		cmd = strings.TrimSpace(cmd)

		argv := strings.Split(cmd, " ")

		switch argv[0] {

		case "exit":
			exit(argv)
		case "echo":
			echo(argv)
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}

	}
}
