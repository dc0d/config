package config

import (
	"os"
	"path/filepath"
)

//-----------------------------------------------------------------------------

/*
RelativeResource returns the file path

	1 - searchs inside dirs, in the same order of dirs
	2 - searchs inside working directory
	3 - searchs inside app directory

otherwise returns os.ErrNotExist error
*/
func RelativeResource(fileName string, dirs ...string) (fp string, fe error) {
	if fp, fe = searchDirs(fileName, dirs...); fe == nil {
		return
	}
	if fp, fe = searchWD(fileName); fe == nil {
		return
	}
	if fp, fe = searchAppDir(fileName); fe == nil {
		return
	}
	fe = os.ErrNotExist
	return
}

func searchDirs(fileName string, dirs ...string) (fp string, fe error) {
	for _, v := range dirs {
		d, err := filepath.Abs(v)
		if err != nil {
			continue
		}
		p := filepath.Join(d, fileName)
		if _, fe = os.Stat(p); fe == nil {
			fp = p
			return
		}
	}
	fe = os.ErrNotExist
	return
}

func searchAppDir(fileName string) (fp string, fe error) {
	var (
		appDir string
	)
	appDir, fe = os.Executable()
	if fe != nil {
		return
	}
	appDir, fe = filepath.Abs(filepath.Dir(appDir))
	if fe != nil {
		return
	}
	p := filepath.Join(appDir, fileName)
	if _, fe = os.Stat(p); fe == nil {
		fp = p
		return
	}
	return
}

func searchWD(fileName string) (fp string, fe error) {
	var (
		wd string
	)
	wd, fe = os.Getwd()
	if fe != nil {
		return
	}
	p := filepath.Join(wd, fileName)
	if _, fe = os.Stat(p); fe == nil {
		fp = p
		return
	}
	return
}

//-----------------------------------------------------------------------------
