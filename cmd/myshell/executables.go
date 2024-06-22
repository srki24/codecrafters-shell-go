package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func getExecutablePath(cmd string) (path string, ok bool) {

	paths := filepath.SplitList(os.Getenv("PATH"))
	for _, p := range paths {
		fp := filepath.Join(p, cmd)

		_, err := exec.LookPath(fp)
		if errors.Is(err, exec.ErrDot) {
			err = nil
		}

		if err == nil {
			path = fp
			ok = true
			break
		}
	}

	return
}

func executeCmd(cmd string) {

	c := exec.Command(cmd)
	if errors.Is(c.Err, exec.ErrDot) {
		c.Err = nil
	}

	if err := c.Run(); err != nil {
		log.Fatal(err)
	}

}
