package commands

import (
	"testing"

	"github.com/KanhaiyaKumarGupta/inito_Assignment/commands"
	"github.com/stretchr/testify/assert"
)

func setupFileSystemWithFiles() *commands.FileSystem {
	fs := commands.NewFileSystem()

	testDir := commands.NewDirectory("test", fs.Root)
	fs.Root.SubDirs["test"] = testDir

	file1 := &commands.File{Name: "file1.txt", Content: "Hello World"}
	file2 := &commands.File{Name: "file2.txt", Content: "Hello Go"}
	testDir.Files["file1.txt"] = file1
	testDir.Files["file2.txt"] = file2

	return fs
}

func TestGrepFoundMatches(t *testing.T) {
	fs := setupFileSystemWithFiles()

	results := fs.Grep("Hello", "/test")
	assert.Contains(t, results, "Match found in file 'file1.txt': Hello World", "Grep should find a match in 'file1.txt'")
	assert.Contains(t, results, "Match found in file 'file2.txt': Hello Go", "Grep should find a match in 'file2.txt'")
}

func TestGrepNoMatch(t *testing.T) {
	fs := setupFileSystemWithFiles()

	results := fs.Grep("Python", "/test")
	assert.Empty(t, results, "Grep should return an empty slice if no match is found")
}

func TestGrepNonExistentDirectory(t *testing.T) {
	fs := commands.NewFileSystem()
	results := fs.Grep("Hello", "/nonexistent")
	assert.Contains(t, results, "Directory not found: /nonexistent", "Grep on a non-existent directory should indicate directory not found")
}
