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

		functions = getAllFunctions()

		if f, ok := getFunction(argv[0]); ok {
			f(argv[1:]...)
		} else if fp, ok := getExecutablePath(argv[0]); ok {
			executeCmd(fp, argv[1:])
		} else {
			fmt.Printf("%s: command not found\n", cmd)
		}

	}
}
