package main

import (
	"github.com/juxuny/yc/cmd/yc/model"
	"github.com/juxuny/yc/cmd/yc/service"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{}
)

func main() {
	rootCmd.AddCommand(service.Command)
	rootCmd.AddCommand(model.Command)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
