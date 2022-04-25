package service

import (
	"github.com/juxuny/yc/cmd"
	"github.com/juxuny/yc/services"
	"github.com/juxuny/yc/utils"
	"github.com/juxuny/yc/utils/template"
	"github.com/spf13/cobra"
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
)

type UpdateCommand struct {
	WorkDir string
}

func (t *UpdateCommand) Prepare(cmd *cobra.Command) {

}

func (t *UpdateCommand) InitFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&t.WorkDir, "work-dir", "w", "", "working dir")
}

func (t *UpdateCommand) Run() {
	if t.WorkDir == "" {
		if w, err := os.Getwd(); err != nil {
			log.Fatal(err)
		} else {
			t.WorkDir = w
		}
	}

	serviceName := t.getServiceName()
	log.Println("service name: ", serviceName)
	service := services.NewServiceEntity(serviceName)
	if err := t.initEnvConfig(service); err != nil {
		log.Fatal(err)
	}
}

func (t *UpdateCommand) getServiceName() string {
	serviceName := ""
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

func (t *UpdateCommand) initEnvConfig(service services.ServiceEntity) error {
	configDir := path.Join(t.WorkDir, "config")
	if err := utils.TouchDir(configDir, 0755); err != nil {
		log.Fatal(err)
	}
	if err := template.RunEmbedFile(templateFs, envConfigFileName, path.Join(t.WorkDir, "config", "env.go"), service); err != nil {
		log.Fatal("create env.go failed:", err)
	}
	return nil
}

func init() {
	builder := cmd.NewCommandBuilder("update", &UpdateCommand{})
	Command.AddCommand(builder.Build())
}
