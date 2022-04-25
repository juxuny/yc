package cmd

import (
	"log"
	"os/exec"
)

func Exec(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	log.Println(cmd.String())
	return cmd.Run()
}
