package service

import (
	"embed"
	"github.com/spf13/cobra"
)

//go:embed template
var templateFs embed.FS

var Service = &cobra.Command{
	Use: "service",
}

func init() {
	Service.AddCommand(createCommand)
}
