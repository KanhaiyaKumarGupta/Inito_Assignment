package commands

import (
	"fmt"
	"strings"
)

func (fs *FileSystem) Mkdir(path string) error {
	parts := strings.Split(path, "/")
	current := fs.Current

	for i, part := range parts {
		if part == "" || part == "." {
			if i == 0 {
				current = fs.Root
			}
			continue
		}

		if part == ".." {
			if current.Parent != nil {
				current = current.Parent
			}
			continue
		}

		if _, exists := current.SubDirs[part]; !exists {
			if i == len(parts)-1 {
				current.SubDirs[part] = NewDirectory(part, current)
			} else {
				return fmt.Errorf("cannot create directory '%s': parent directory does not exist", strings.Join(parts[:i+1], "/"))
			}
		} else if i == len(parts)-1 {

			return fmt.Errorf("directory '%s' already exists", path)
		}

		current = current.SubDirs[part]
	}
	return nil
}
