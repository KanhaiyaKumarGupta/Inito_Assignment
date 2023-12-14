package commands

import (
	"fmt"
	"strings"
)

func (fs *FileSystem) Touch(filePath string) error {
	parts := strings.Split(filePath, "/")
	fileName := parts[len(parts)-1]
	dirPath := strings.Join(parts[:len(parts)-1], "/")

	targetDir := fs.GetDir(dirPath)
	if targetDir == nil {
		return fmt.Errorf("directory not found: %s", dirPath)
	}

	if _, exists := targetDir.Files[fileName]; exists {

		return fmt.Errorf("file already exists: %s", fileName)
	}

	targetDir.Files[fileName] = &File{Name: fileName, Content: ""}
	return nil
}
