package utils

import (
	"errors"
	"github.com/jszwec/csvutil"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func SaveAsCSV[D any](file string, data D) error {
	if _, err := os.Stat(filepath.Dir(file)); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(filepath.Dir(file), os.ModePerm)
		if err != nil {
			return err
		}
	}
	marshal, err := csvutil.Marshal(data)
	if err != nil {
		return err
	}
	permissions := 0644 // or whatever you need
	err = ioutil.WriteFile(file, marshal, fs.FileMode(permissions))
	if err != nil {
		return err
	}
	return nil
}
