package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func getExecutablePath(name string) (path string, ok bool) {

	paths := filepath.SplitList(os.Getenv("PATH"))
	for _, p := range paths {
		fp := filepath.Join(p, name)

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

func executeCmd(name string, args []string) {

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

}
