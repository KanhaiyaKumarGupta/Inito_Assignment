package commands

import (
	"testing"

	"github.com/KanhaiyaKumarGupta/inito_Assignment/commands"
	"github.com/stretchr/testify/assert"
)

func setupFileSystemForRmTest() *commands.FileSystem {
	fs := commands.NewFileSystem()
	fs.Mkdir("/test")
	fs.Touch("/test/file.txt")
	fs.Mkdir("/test/dir")
	return fs
}

func TestRmFile(t *testing.T) {
	fs := setupFileSystemForRmTest()

	err := fs.Rm("/test/file.txt")
	assert.NoError(t, err, "Removing file should not produce an error")

	_, exists := fs.GetDir("/test").Files["file.txt"]
	assert.False(t, exists, "File should be removed from the directory")
}

func TestRmDirectory(t *testing.T) {
	fs := setupFileSystemForRmTest()

	err := fs.Rm("/test/dir")
	assert.NoError(t, err, "Removing directory should not produce an error")

	_, exists := fs.GetDir("/test").SubDirs["dir"]
	assert.False(t, exists, "Directory should be removed from the directory")
}

func TestRmNonExistentItem(t *testing.T) {
	fs := setupFileSystemForRmTest()

	err := fs.Rm("/test/nonexistent.txt")
	assert.Error(t, err, "Attempting to remove a non-existent item should return an error")
}
