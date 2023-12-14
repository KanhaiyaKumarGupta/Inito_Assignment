package commands

import (
	"fmt"
	"strings"
)

func (fs *FileSystem) Echo(text, filePath string) error {
	parts := strings.Split(filePath, "/")
	fileName := parts[len(parts)-1]
	dirPath := strings.Join(parts[:len(parts)-1], "/")

	targetDir := fs.GetDir(dirPath)
	if targetDir == nil {
		return fmt.Errorf("directory not found: %s", dirPath)
	}


	file, exists := targetDir.Files[fileName]
	if !exists {
		file = &File{Name: fileName}
		targetDir.Files[fileName] = file
	}
	file.Content += text 
	return nil
}
