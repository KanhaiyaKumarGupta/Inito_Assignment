package commands

import (
	"fmt"
	"strings"
)

func (fs *FileSystem) Rm(path string) error {
	dirPath := strings.Join(strings.Split(path, "/")[:len(strings.Split(path, "/"))-1], "/")
	itemName := strings.Split(path, "/")[len(strings.Split(path, "/"))-1]

	targetDir := fs.GetDir(dirPath)
	if targetDir == nil {
		return fmt.Errorf("directory not found: %s", dirPath)
	}

	if _, exists := targetDir.Files[itemName]; exists {
		delete(targetDir.Files, itemName)
		return nil
	}

	if _, exists := targetDir.SubDirs[itemName]; exists {
		delete(targetDir.SubDirs, itemName)
		return nil
	}

	return fmt.Errorf("item not found: %s", itemName)
}
