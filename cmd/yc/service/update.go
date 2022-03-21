package service

import "github.com/spf13/cobra"

var updateCommand = &cobra.Command{
	Use: "update",
	Run: update,
}

func update(cmd *cobra.Command, args []string) {

}
