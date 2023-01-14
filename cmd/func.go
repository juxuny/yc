package cmd

import (
	"github.com/juxuny/yc/errors"
	"os"
	"strings"
)

func CheckIfCommandExists(command string) (ok bool, err error) {
	value := os.Getenv("PATH")
	pathList := strings.Split(value, ":")
	if len(pathList) == 0 {
		return false, nil
	}
	for _, p := range pathList {
		files, err := os.ReadDir(p)
		if err != nil {
			return false, errors.SystemError.FsError.Wrap(err)
		}
		for _, f := range files {
			if f.Name() == command {
				return true, nil
			}
		}
	}
	return false, nil
}
