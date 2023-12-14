package commands

import (
	"fmt"
	"strings"
)

func (fs *FileSystem) Mv(srcPath, destPath string) error {
	srcDirPath := strings.Join(strings.Split(srcPath, "/")[:len(strings.Split(srcPath, "/"))-1], "/")
	srcItemName := strings.Split(srcPath, "/")[len(strings.Split(srcPath, "/"))-1]

	destDirPath := strings.Join(strings.Split(destPath, "/")[:len(strings.Split(destPath, "/"))-1], "/")
	destItemName := strings.Split(destPath, "/")[len(strings.Split(destPath, "/"))-1]

	srcDir := fs.GetDir(srcDirPath)
	destDir := fs.GetDir(destDirPath)
	if srcDir == nil || destDir == nil {
		return fmt.Errorf("invalid path")
	}

	if _, exists := destDir.Files[destItemName]; exists {
		return fmt.Errorf("file already exists at destination: %s", destItemName)
	}
	if _, exists := destDir.SubDirs[destItemName]; exists {
		return fmt.Errorf("directory already exists at destination: %s", destItemName)
	}

	if file, exists := srcDir.Files[srcItemName]; exists {
		delete(srcDir.Files, srcItemName)
		destDir.Files[destItemName] = file
		return nil
	}

	if dir, exists := srcDir.SubDirs[srcItemName]; exists {
		delete(srcDir.SubDirs, srcItemName)
		destDir.SubDirs[destItemName] = dir
		dir.Parent = destDir
		return nil
	}

	return fmt.Errorf("item not found: %s", srcItemName)
}
