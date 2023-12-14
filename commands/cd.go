package commands

import (
	"fmt"
	"strings"
)

func (fs *FileSystem) Cd(path string) error {
	if path == "/" || path == "~" {
		fs.Current = fs.Root
		return nil
	}

	var target *Directory
	if strings.HasPrefix(path, "/") {
		target = fs.Root
		path = strings.TrimPrefix(path, "/")
	} else {
		target = fs.Current
	}

	parts := strings.Split(path, "/")
	for _, part := range parts {
		if part == "" || part == "." {
			continue
		}

		if part == ".." {
			if target.Parent != nil {
				target = target.Parent
			}
			continue
		}

		next, ok := target.SubDirs[part]
		if !ok {
			return fmt.Errorf("directory not found: %s", part)
		}
		target = next
	}

	fs.Current = target
	return nil
}
