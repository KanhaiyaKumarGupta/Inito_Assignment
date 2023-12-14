package commands

import (
	"fmt"
	"strings"
)

type File struct {
	Name    string
	Content string
}

type Directory struct {
	Name    string
	Files   map[string]*File
	SubDirs map[string]*Directory
	Parent  *Directory
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		Name:    name,
		Files:   make(map[string]*File),
		SubDirs: make(map[string]*Directory),
		Parent:  parent,
	}
}

type FileSystem struct {
	Root    *Directory
	Current *Directory
}

func NewFileSystem() *FileSystem {
	root := NewDirectory("/", nil)
	return &FileSystem{
		Root:    root,
		Current: root,
	}
}

func (fs *FileSystem) GetDir(path string) *Directory {
	var current *Directory
	if strings.HasPrefix(path, "/") {
		current = fs.Root
		path = strings.TrimPrefix(path, "/")
	} else {
		current = fs.Current
	}

	parts := strings.Split(path, "/")
	for _, part := range parts {
		if part == "" || part == "." {
			continue
		}
		if part == ".." {
			if current.Parent != nil {
				current = current.Parent
			}
			continue
		}
		next, ok := current.SubDirs[part]
		if !ok {
			fmt.Println("Directory not found:", part)
			return nil
		}
		current = next
	}
	return current
}
func (fs *FileSystem) _parsePath(path string) (*Directory, string) {
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return nil, ""
	}

	dirPath := strings.Join(parts[:len(parts)-1], "/")
	itemName := parts[len(parts)-1]

	dir := fs.GetDir(dirPath)
	return dir, itemName
}
