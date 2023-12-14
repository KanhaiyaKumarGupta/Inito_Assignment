package commands

import (
	"fmt"
	"strings"
)

func (fs *FileSystem) Cat(filePath string) (string, error) {

	parts := strings.Split(filePath, "/")
	dirPath := strings.Join(parts[:len(parts)-1], "/")
	fileName := parts[len(parts)-1]

	targetDir := fs.GetDir(dirPath)
	if targetDir == nil {
		return "", fmt.Errorf("directory not found: %s", dirPath)
	}

	file, exists := targetDir.Files[fileName]
	if !exists {
		return "", fmt.Errorf("file not found: %s", fileName)
	}

	return file.Content, nil
}
