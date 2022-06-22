package client

import (
	"github.com/juxuny/yc/cmd"
	"github.com/juxuny/yc/utils"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

type GenCommand struct {
	Go       string
	ReactTs  string
	WorkDir  string
	Internal bool
}

func (t *GenCommand) Prepare(cmd *cobra.Command) {
}

func (t *GenCommand) InitFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&t.Go, "go", "", "go sdk output directory")
	cmd.PersistentFlags().StringVar(&t.ReactTs, "react-ts", "", "react ts sdk output directory")
	cmd.PersistentFlags().StringVarP(&t.WorkDir, "work-dir", "w", ".", "working directory")
	cmd.PersistentFlags().BoolVar(&t.Internal, "internal", false, "auto gen internal RPC method")
}

func (t *GenCommand) BeforeRun(cmd *cobra.Command) {
	if t.Go == "" && t.ReactTs == "" {
		log.Fatal("missing args: --go, --react-ts")
	}
}

func (t *GenCommand) Run() {
	if t.Go != "" {
		t.genGo()
	}
	if t.ReactTs != "" {
		t.genReactTs()
	}
}

func (t *GenCommand) getServiceName() string {
	serviceName := ""
	if _, err := os.Stat(t.WorkDir); os.IsNotExist(err) {
		log.Fatal(err)
	}
	if err := filepath.Walk(t.WorkDir, func(filePath string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if path.Ext(filePath) == ".proto" {
			serviceName = path.Base(filePath)
			return nil
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return utils.StringHelper.TrimSubSequenceRight(serviceName, ".proto")
}

func init() {
	genCommand := &GenCommand{}
	builder := cmd.NewCommandBuilder("gen", genCommand)
	Command.AddCommand(builder.Build())
}
