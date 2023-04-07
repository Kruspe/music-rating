package setup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type CommandInBackground struct {
	Command string
	Args    []string
	Dir     string
}

func RunCommandInBackground(name string, c CommandInBackground) error {
	_, f, _, _ := runtime.Caller(0)
	rootDir := filepath.Join(filepath.Dir(f), "../..")
	log, err := os.Create(filepath.Join(rootDir, fmt.Sprintf("logs/%s.log", name)))
	if err != nil {
		return err
	}

	cmd := exec.Command(c.Command, c.Args...)
	cmd.Stdout = log
	cmd.Stderr = log
	cmd.Dir = filepath.Join(rootDir, c.Dir)

	return cmd.Start()
}
