package gen

import (
	"github.com/juxuny/yc/cmd"
	"github.com/spf13/cobra"
)

type ModelCommand struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Prefix   string
	Service  string
}

func (t *ModelCommand) Prepare(cmd *cobra.Command) {
}

func (t *ModelCommand) InitFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&t.Name, "name", "n", "", "schema name")
	cmd.PersistentFlags().StringVarP(&t.Host, "host", "h", "127.0.0.1", "database host")
	cmd.PersistentFlags().IntVar(&t.Port, "port", 3306, "database port")
	cmd.PersistentFlags().StringVarP(&t.User, "user", "u", "", "user")
	cmd.PersistentFlags().StringVarP(&t.Password, "password", "p", "", "password")
}

func (t *ModelCommand) Run() {
}

func init() {
	Command.AddCommand(cmd.NewCommandBuilder("model", &ModelCommand{}).Build())
}
