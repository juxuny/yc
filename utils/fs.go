package utils

import (
	"github.com/juxuny/yc/errors"
	"os"
)

func TouchDir(dir string, perm os.FileMode) error {
	stat, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(dir, perm)
		}
	}
	if !stat.IsDir() {
		return errors.SystemError.FsIsNotDir
	}
	return nil
}

func IsFileOrDirExists(p string) bool {
	_, stat := os.Stat(p)
	return stat == nil
}

func IsFileOrDirNotExists(p string) bool {
	_, err := os.Stat(p)
	return os.IsNotExist(err)
}
