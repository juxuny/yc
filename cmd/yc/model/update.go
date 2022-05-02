package model

import (
	"github.com/juxuny/yc/cmd"
	"github.com/spf13/cobra"
)

type UpdateCommand struct {
	WorkDir string
}

func (t *UpdateCommand) Prepare(cmd *cobra.Command) {

}

func (t *UpdateCommand) InitFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&t.WorkDir, "work-dir", "w", "", "working dir")
}

func (t *UpdateCommand) BeforeRun(cmd *cobra.Command) {

}

func (t *UpdateCommand) Run() {

}

func init() {
	builder := cmd.NewCommandBuilder("update", &UpdateCommand{})
	Command.AddCommand(builder.Build())
}
