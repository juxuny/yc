//go.mod:build linux
//go:build linux
// +build linux

package service

import (
	"github.com/juxuny/yc/utils"
	"log"
	"path"
)

func installProtoc(outputDir string) error {
	archiveBaseName := path.Base(protocBinaryArchiveFileNameLinux)
	archiveOutputDirName := utils.StringHelper.TrimSubSequenceRight(archiveBaseName, path.Ext(archiveBaseName))
	outputDir = path.Join(outputDir, archiveOutputDirName)
	if err := utils.TouchDir(outputDir, 0776); err != nil {
		log.Fatalln(err)
	}
	archiveFileName := path.Join(outputDir, archiveBaseName)
	err := extractEmbedFileName(protocBinaryArchiveFileNameLinux, archiveFileName)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
