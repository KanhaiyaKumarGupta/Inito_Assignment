package commands

import (
	"testing"

	"github.com/KanhaiyaKumarGupta/inito_Assignment/commands"
	"github.com/stretchr/testify/assert"
)

func setupFileSystemForCpTest() *commands.FileSystem {
	fs := commands.NewFileSystem()
	fs.Mkdir("/src")
	fs.Mkdir("/dest")
	fs.Touch("/src/file.txt")
	fs.Mkdir("/src/dir")
	return fs
}

func TestCpFile(t *testing.T) {
	fs := setupFileSystemForCpTest()
	err := fs.Cp("/src/file.txt", "/dest/file.txt")
	assert.NoError(t, err, "Copying file should not produce an error")

	_, existsInSrc := fs.GetDir("/src").Files["file.txt"]
	_, existsInDest := fs.GetDir("/dest").Files["file.txt"]
	assert.True(t, existsInSrc, "File should still exist in the source directory")
	assert.True(t, existsInDest, "File should be copied to the destination directory")
}

func TestCpDirectory(t *testing.T) {
	fs := setupFileSystemForCpTest()

	err := fs.Cp("/src/dir", "/dest/dir")
	assert.NoError(t, err, "Copying directory should not produce an error")

	_, existsInSrc := fs.GetDir("/src").SubDirs["dir"]
	_, existsInDest := fs.GetDir("/dest").SubDirs["dir"]
	assert.True(t, existsInSrc, "Directory should still exist in the source directory")
	assert.True(t, existsInDest, "Directory should be copied to the destination directory")
}

func TestCpNonExistentItem(t *testing.T) {
	fs := setupFileSystemForCpTest()

	err := fs.Cp("/src/nonexistent.txt", "/dest/nonexistent.txt")
	assert.Error(t, err, "Attempting to copy a non-existent item should return an error")
}
