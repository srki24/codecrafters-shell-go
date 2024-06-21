package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func getExecutablePath(cmd string) (path string, ok bool) {

	ok = false

	paths := filepath.SplitList(os.Getenv("PATH"))
	for _, p := range paths {
		fp := filepath.Join(p, cmd)

		if _, err := exec.LookPath(fp); err == nil {
			path = p
			ok = true
			break
		}
	}

	return
}

func executeCmd(cmd string) {

	c := exec.Command(cmd)
	if err := c.Run(); err != nil {
		log.Fatal(err)
	}

}
