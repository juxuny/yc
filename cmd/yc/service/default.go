package service

import (
	"embed"
	"github.com/spf13/cobra"
)

//go:embed template
var templateFs embed.FS

var Command = &cobra.Command{
	Use: "service",
}
