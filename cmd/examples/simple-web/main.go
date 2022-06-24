package main

import (
	"github.com/juxuny/yc/cmd"
	"github.com/spf13/cobra"
	"log"
)

type serviceCommand struct {
	port int
}

func (t *serviceCommand) Prepare(cmd *cobra.Command) {
}

func (t *serviceCommand) InitFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().IntVar(&t.port, "port", 9000, "http port")
}

func (t *serviceCommand) BeforeRun(cmd *cobra.Command) {
	if t.port < 1024 {
		log.Fatal("don't use port less than 1024")
	}
}

func (t *serviceCommand) Run() {

}

func main() {
	rootCommand := cmd.NewCommandBuilder("web", &serviceCommand{}).Build()
	if err := rootCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
