package cmd

import "github.com/spf13/cobra"

type commandBuilder struct {
	Name   string
	Runner Runner
	Desc   string
}

func NewCommandBuilder(name string, runner Runner) *commandBuilder {
	return &commandBuilder{
		Name:   name,
		Runner: runner,
	}
}

func (t *commandBuilder) Build() *cobra.Command {
	cmd := &cobra.Command{
		Use: t.Name,
		Run: func(cmd *cobra.Command, args []string) {
			t.Runner.Run()
		},
	}
	t.Runner.Prepare(cmd)
	t.Runner.InitFlag(cmd)
	return cmd
}

type Runner interface {
	Prepare(cmd *cobra.Command)
	InitFlag(cmd *cobra.Command)
	Run()
}
