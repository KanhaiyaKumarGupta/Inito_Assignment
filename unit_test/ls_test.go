package commands

import (
	"testing"

	"github.com/KanhaiyaKumarGupta/inito_Assignment/commands"
	"github.com/stretchr/testify/assert"
)

func TestLs(t *testing.T) {
	fs := commands.NewFileSystem()
	fs.Mkdir("/test")
	fs.Touch("/test/file1.txt")
	fs.Mkdir("/test/subdir")
	contents := fs.Ls("/test")
	assert.Contains(t, contents, "file1.txt", "Ls should list 'file1.txt'")
	assert.Contains(t, contents, "subdir/", "Ls should list 'subdir/'")

	rootContents := fs.Ls("/")
	assert.Contains(t, rootContents, "test/", "Ls should list 'test/' in the root directory")
}

func TestLsNonExistentDirectory(t *testing.T) {
	fs := commands.NewFileSystem()
	contents := fs.Ls("/nonexistent")
	assert.Contains(t, contents, "Directory not found: /nonexistent", "Ls on a non-existent directory should indicate directory not found")
}
