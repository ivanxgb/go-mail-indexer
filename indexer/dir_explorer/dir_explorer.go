package dir_explorer

import (
	"os"
	"path/filepath"
)

const (
	extensionFile = "."
)

// GetFilesPath receives a directory path and returns a slice of strings with
// the path of all the files inside the directory.
func GetFilesPath(dirPath string) ([]string, error) {
	var filesPath []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// If the path is not a directory or the file extension is valid, is added to the slice.
		if !info.IsDir() || filepath.Ext(path) == extensionFile {
			filesPath = append(filesPath, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return filesPath, nil
}
