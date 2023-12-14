package commands

import (
	"fmt"
	"strings"
)

func (fs *FileSystem) Grep(pattern, path string) []string {
	target := fs.GetDir(path)
	if target == nil {
		return []string{"Directory not found: " + path}
	}

	var matches []string
	for _, file := range target.Files {
		if strings.Contains(file.Content, pattern) {
			match := fmt.Sprintf("Match found in file '%s': %s", file.Name, file.Content)
			matches = append(matches, match)
		}
	}
	return matches
}
