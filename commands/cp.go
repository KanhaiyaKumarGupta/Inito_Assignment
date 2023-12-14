package commands

import (
	"fmt"
)

func copyDirectory(src *Directory, parent *Directory) *Directory {
	newDir := NewDirectory(src.Name, parent)
	for name, file := range src.Files {
		newFile := &File{Name: file.Name, Content: file.Content}
		newDir.Files[name] = newFile
	}
	for name, subDir := range src.SubDirs {
		newDir.SubDirs[name] = copyDirectory(subDir, newDir)
	}
	return newDir
}

func (fs *FileSystem) Cp(srcPath, destPath string) error {
	srcDir, srcName := fs._parsePath(srcPath)
	destDir, _ := fs._parsePath(destPath)

	if srcDir == nil || destDir == nil {
		return fmt.Errorf("invalid path")
	}

	if _, exists := destDir.Files[srcName]; exists {
		return fmt.Errorf("file already exists at destination: %s", srcName)
	}
	if _, exists := destDir.SubDirs[srcName]; exists {
		return fmt.Errorf("directory already exists at target: %s", srcName)
	}

	if file, exists := srcDir.Files[srcName]; exists {
		newFile := &File{Name: file.Name, Content: file.Content}
		destDir.Files[srcName] = newFile
		return nil
	}

	if dir, exists := srcDir.SubDirs[srcName]; exists {
		destDir.SubDirs[srcName] = copyDirectory(dir, destDir)
		return nil
	}

	return fmt.Errorf("item not found: %s", srcName)
}
