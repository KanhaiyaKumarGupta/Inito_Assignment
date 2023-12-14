package commands

import (
	"testing"

	"github.com/KanhaiyaKumarGupta/inito_Assignment/commands"
	"github.com/stretchr/testify/assert"
)

func TestMvFile(t *testing.T) {
	fs := commands.NewFileSystem()
	fs.Mkdir("/src")
	fs.Mkdir("/dest")
	fs.Touch("/src/testfile.txt")

	err := fs.Mv("/src/testfile.txt", "/dest/testfile.txt")
	assert.NoError(t, err, "Moving file should not produce an error")

	_, exists := fs.GetDir("/src").Files["testfile.txt"]
	assert.False(t, exists, "File should not exist in the source directory after moving")

	_, exists = fs.GetDir("/dest").Files["testfile.txt"]
	assert.True(t, exists, "File should exist in the destination directory after moving")
}

func TestMvDirectory(t *testing.T) {
	fs := commands.NewFileSystem()
	fs.Mkdir("/srcdir")
	fs.Mkdir("/destdir")

	err := fs.Mv("/srcdir", "/destdir/srcdir")
	assert.NoError(t, err, "Moving directory should not produce an error")

	_, exists := fs.GetDir("/").SubDirs["srcdir"]
	assert.False(t, exists, "Directory should not exist at the original location after moving")

	_, exists = fs.GetDir("/destdir").SubDirs["srcdir"]
	assert.True(t, exists, "Directory should exist at the new location after moving")
}

func TestMvNonExistentItem(t *testing.T) {
	fs := commands.NewFileSystem()

	err := fs.Mv("/nonexistent.txt", "/dest/nonexistent.txt")

	assert.Error(t, err, "Attempting to move a non-existent item should return an error")
}
