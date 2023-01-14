package service

import (
	"github.com/juxuny/yc/cmd"
	"github.com/juxuny/yc/utils"
	"github.com/juxuny/yc/utils/template"
	"log"
	"os"
	"path"
)

func extractEmbedFileName(srcEmbedFileName string, outputFileName string) error {
	err := template.SaveEmbedFileAs(templateFs, srcEmbedFileName, outputFileName)
	if err != nil {
		log.Fatalln(err)
	}
	extractDir := path.Dir(outputFileName)
	err = cmd.Exec("unzip", "-d", extractDir, outputFileName)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

func prepareGrpc() {
	ok, err := cmd.CheckIfCommandExists("protoc")
	if err != nil {
		log.Fatalln(err)
	}
	if !ok {
		log.Println("init protoc")
		gopath, found := os.LookupEnv("GOPATH")
		if !found {
			log.Fatalln("not found GOPATH in environment variable list")
		}
		gopathBin := path.Join(gopath, "bin")
		err = utils.TouchDir(gopathBin, 0666)
		if err != nil {
			log.Fatalln(err)
		}
		err = installProtoc(gopathBin)
		if err != nil {
			log.Fatalln(err)
		}
	}
	ok, err = cmd.CheckIfCommandExists("protoc-gen-go")
	if err != nil {
		log.Fatalln(err)
	}
	if !ok {
		err = cmd.Exec("go", "install", "google.golang.org/protobuf/cmd/protoc-gen-go@v1.28")
		if err != nil {
			log.Fatalln(err)
		}
	}
	ok, err = cmd.CheckIfCommandExists("protoc-gen-go-grpc")
	if err != nil {
		log.Fatalln(err)
	}
	if !ok {
		err = cmd.Exec("go", "install", "google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2")
		if err != nil {
			log.Fatalln(err)
		}
	}
}
