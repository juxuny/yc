package service

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var createCommand = &cobra.Command{
	Use: "create",
	Run: create,
}

func create(cmd *cobra.Command, args []string) {
	fis, err := templateFs.ReadDir("template")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range fis {
		fmt.Println(item.Name())
	}
}
