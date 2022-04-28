package gen

import (
	"github.com/juxuny/yc/cmd"
	"github.com/spf13/cobra"
)

type RpcCommand struct {
}

func (r RpcCommand) BeforeRun(cmd *cobra.Command) {
	panic("implement me")
}

func (r RpcCommand) Prepare(cmd *cobra.Command) {
}

func (r RpcCommand) InitFlag(cmd *cobra.Command) {
}

func (r RpcCommand) Run() {
}

func init() {
	Command.AddCommand(cmd.NewCommandBuilder("rpc", &RpcCommand{}).Build())
}
