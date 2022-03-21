package main

import (
	"github.com/juxuny/yc/cmd/yc/service"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{}
)

func main() {
	rootCmd.AddCommand(service.Service)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
