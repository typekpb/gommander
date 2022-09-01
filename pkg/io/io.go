package io

import (
	"os"
)

func OSReadDir(root string) ([]os.FileInfo, error) {
	var files []os.FileInfo
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file)
	}
	return files, nil
}