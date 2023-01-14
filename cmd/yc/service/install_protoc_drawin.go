//go.mod:build drawin
//go:build drawin
// +build drawin

package service

func installProtoc(outputDir string) error {
	archiveBaseName := path.Base(protocBinaryArchiveFileNameOSX)
	archiveOutputDirName := utils.StringHelper.TrimSubSequenceRight(archiveBaseName, path.Ext(archiveBaseName))
	outputDir = path.Join(outputDir, archiveOutputDirName)
	if err := utils.TouchDir(outputDir, 0776); err != nil {
		log.Fatalln(err)
	}
	archiveFileName := path.Join(outputDir, archiveBaseName)
	err := extractEmbedFileName(protocBinaryArchiveFileNameOSX, archiveFileName)
	if err != nil {
		log.Fatalln(err)
	}
	gopath, found := os.LookupEnv("GOPATH")
	if !found {
		log.Fatalln("missing GOPATH")
	}
	protocBinFileName := path.Join(gopath, "bin", "protoc")
	if err := cmd.Exec("ln", "-s", path.Join(outputDir, "bin", "protoc"), protocBinFileName); err != nil {
		log.Fatalln(err)
	}
	return nil
}
