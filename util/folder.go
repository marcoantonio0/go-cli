package util

import (
	"os"
	"path/filepath"
)

func CreateFolder(folderName string) error {
	err := os.Mkdir(folderName, 0775)
	if err != nil {
		return nil
	}
	return nil
}

func CreateFolderWithParent(parentDir string, folderName string) error {
	dir := filepath.Join(parentDir, folderName)
	err := os.MkdirAll(dir, 0775)
	if err != nil {
		return nil
	}
	return nil
}
