package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateFile creates a file along with the complete path.
func CreateFile(filePath string) error {
	var err error
	if _, err = os.Stat(filePath); err != nil {
		dirPath := filepath.Dir(filePath)
		if dirPath != "." {
			// create the directory path if it doesn't exist
			if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
				return fmt.Errorf("error creating csv %s: %v", filePath, err)
			}
		}
		_, err = os.Create(filePath)
		if err != nil {
			return fmt.Errorf("error creating file: %v", err)
		}

	}
	return nil
}
