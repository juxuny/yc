package main

import (
	"github.com/juxuny/yc/cmd"
	"github.com/spf13/cobra"
	"log"
)

var (
	rootCmd = &cobra.Command{}
)

func main() {
	rootCmd.AddCommand(cmd.NewCommandBuilder("get", &getCommand{}).Build())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
