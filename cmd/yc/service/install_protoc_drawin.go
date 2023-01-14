//go.mod:build drawin
//go:build drawin
// +build drawin

package service

func installProtoc(outputDir string) error {
	archiveBaseName := path.Base(protocBinaryArchiveFileNameOSX)
	archiveFileName := path.Join(outputDir, archiveBaseName)
	err := extractEmbedFileName(protocBinaryArchiveFileNameOSX, archiveFileName)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
