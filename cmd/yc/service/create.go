package service

import (
	"github.com/juxuny/yc"
	"github.com/juxuny/yc/cmd"
	"github.com/juxuny/yc/services"
	"github.com/juxuny/yc/utils"
	"github.com/juxuny/yc/utils/template"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path"
)

type CreateCommand struct {
	Name      string
	WorkDir   string
	OutputDir string
}

func (t *CreateCommand) BeforeRun(cmd *cobra.Command) {
	prepareGrpc()
}

func (t *CreateCommand) Prepare(cmd *cobra.Command) {
}

func (t *CreateCommand) InitFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&t.Name, "name", "n", "", "service name")
	cmd.PersistentFlags().StringVarP(&t.WorkDir, "work-dir", "w", "", "working dir")
	cmd.PersistentFlags().StringVarP(&t.OutputDir, "output", "o", "", "output server dir")
}

func (t *CreateCommand) Run() {
	if t.Name == "" {
		log.Fatal("missing --name")
	}
	if t.WorkDir == "" {
		if w, err := os.Getwd(); err != nil {
			log.Fatal(err)
		} else {
			t.WorkDir = w
		}
	}
	service := services.NewServiceEntity(t.Name, yc.Version)
	serviceDir := path.Join(t.WorkDir, service.ServiceDir)
	if t.OutputDir == "" {
		if utils.IsFileOrDirExists(serviceDir) {
			log.Fatalf("service '%s' is exists", service.ServiceDir)
		}
		if err := utils.TouchDir(service.ServiceDir, 0755); err != nil {
			log.Fatal(err)
		}
	} else {
		serviceDir = t.OutputDir
		if utils.IsFileOrDirNotExists(serviceDir) {
			if err := utils.TouchDir(serviceDir, 0666); err != nil {
				log.Println("create service dir failed")
				log.Fatal(err)
			}
		}
	}

	// generate .proto
	protoOutputFile := path.Join(t.WorkDir, service.ServiceDir, service.PackageName+".proto")
	if t.OutputDir != "" {
		protoOutputFile = path.Join(t.OutputDir, service.PackageName+".proto")
	}
	log.Printf("creating proto file: %s", protoOutputFile)
	if err := template.RunEmbedFile(templateFs, protoFileName, protoOutputFile, service); err != nil {
		log.Fatal(err)
	}

	log.Printf("create service finished: %s", service.ServiceDir)
}

func init() {
	createCommand := &CreateCommand{}
	builder := cmd.NewCommandBuilder("create", createCommand)
	Command.AddCommand(builder.Build())
}
