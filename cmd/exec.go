package cmd

import (
	"log"
	"os"
	"os/exec"
)

func Exec(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	log.Println(cmd.String())
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
