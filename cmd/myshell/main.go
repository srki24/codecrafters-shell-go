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

		if argv[0] == "exit" {
			var i int
			if _, err := fmt.Sscan(argv[1], &i); err == nil {
				os.Exit(i)
			} else {
				log.Fatal(err)
			}
		}

		fmt.Printf("%s: command not found\n", cmd)

	}
}
