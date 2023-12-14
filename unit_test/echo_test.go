package commands

import (
	"testing"

	"github.com/KanhaiyaKumarGupta/inito_Assignment/commands"
	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
	fs := commands.NewFileSystem()
	fs.Mkdir("/test")

	err := fs.Echo("Hello World", "/test/newfile.txt")
	assert.NoError(t, err, "Echo should not produce an error when writing to a new file")

	file, exists := fs.GetDir("/test").Files["newfile.txt"]
	assert.True(t, exists, "File should be created in the directory")
	assert.Equal(t, "Hello World", file.Content, "Content of the file should match")

	err = fs.Echo(" More text", "/test/newfile.txt")
	assert.NoError(t, err, "Echo should not produce an error when appending to an existing file")
	assert.Equal(t, "Hello World More text", file.Content, "Content of the file should be appended correctly")
}

func TestEchoNonExistentDirectory(t *testing.T) {
	fs := commands.NewFileSystem()

	err := fs.Echo("Hello World", "/nonexistent/newfile.txt")
	assert.Error(t, err, "Echo should produce an error when the directory does not exist")
}
