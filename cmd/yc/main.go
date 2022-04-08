package main

import (
	"github.com/juxuny/yc/cmd/yc/gen"
	"github.com/juxuny/yc/cmd/yc/service"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{}
)

func main() {
	rootCmd.AddCommand(service.Command)
	rootCmd.AddCommand(gen.Command)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
