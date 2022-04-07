package test

import (
	"path"
	"path/filepath"
	"runtime"
)

func getPath() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return path.Join(filepath.Dir(d), ".")
}

func GetFile(filepath string) string {
	return path.Join(getPath(), filepath)
}
