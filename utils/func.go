package utils

import (
	"fmt"
	"github.com/juxuny/yc/errors"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func GetModuleName(w string) (packageName string, err error) {
	stat, err := os.Stat(w)
	if err != nil {
		return "", errors.SystemError.FsIsNotDir.Wrap(err)
	}
	if !stat.IsDir() {
		return "", errors.SystemError.FsIsNotDir.WithField("path", w)
	}
	stat, err = os.Stat(path.Join(w, "go.mod"))
	if err != nil {
		if os.IsNotExist(err) {
			if len(w) == 1 {
				return "", errors.SystemError.InvalidGoModule
			}
			return GetModuleName(path.Dir(w))
		} else {
			return "", errors.SystemError.InvalidGoModule.Wrap(err)
		}
	}
	return getModuleFromGoModFile(path.Join(w, "go.mod"))
}

func getModuleFromGoModFile(file string) (packageName string, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", errors.SystemError.InvalidGoModule.Wrap(err)
	}
	lines := strings.Split(string(data), "\n")
	for _, l := range lines {
		if strings.Contains(l, "module") {
			packageName = strings.Replace(l, "module", "", 1)
			packageName = strings.TrimSpace(packageName)
			if packageName == "" {
				return "", errors.SystemError.InvalidGoModule.Wrap(fmt.Errorf("package name is empty"))
			}
			return
		}
	}
	return "", errors.SystemError.InvalidGoModule.Wrap(fmt.Errorf("invalid go.mod, not found module definition"))
}

func GetCurrentPackageName(w string) (packageName string, err error) {
	if w == "." {
		w, _ = os.Getwd()
	}
	stat, err := os.Stat(w)
	if err != nil {
		return "", errors.SystemError.FsIsNotDir.Wrap(err)
	}
	if !stat.IsDir() {
		return "", errors.SystemError.FsIsNotDir.WithField("path", w)
	}
	stat, err = os.Stat(path.Join(w, "go.mod"))
	pathSplit := make([]string, 0)
	if err != nil {
		if os.IsNotExist(err) {
			if len(w) == 1 {
				return "", errors.SystemError.InvalidGoModule
			}
			pathSplit = append([]string{path.Base(w)}, pathSplit...)
			packageName, err = GetCurrentPackageName(path.Dir(w))
			packageName = path.Join(append([]string{packageName}, pathSplit...)...)
			return
		} else {
			return "", errors.SystemError.InvalidGoModule.Wrap(err)
		}
	}
	return getModuleFromGoModFile(path.Join(w, "go.mod"))
}
