package ymlconfig

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dc0d/config"
	yaml "gopkg.in/yaml.v2"
)

// New .
func New() config.Loader {
	return config.LoaderFunc(loadYML)
}

func loadYML(ptr interface{}, filePath ...string) error {
	fp, err := checkFilePath(filePath...)
	if err != nil {
		return err
	}
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(content, ptr)
}

func checkFilePath(filePath ...string) (fp string, funcErr error) {
	if len(filePath) > 0 {
		fp = filePath[0]
	}

	if fp == "" {
		var err error
		appName := filepath.Base(os.Args[0])
		appConfName := fmt.Sprintf("%s.ini", appName)
		if fp, err = config.RelativeResource(appConfName); err != nil {
			if fp, err = config.RelativeResource(genericConfName); err != nil {
				return "", err
			}
		}
	}
	if _, err := os.Stat(fp); err != nil {
		return "", err
	}
	return
}

const (
	genericConfName = "app.yml"
)
