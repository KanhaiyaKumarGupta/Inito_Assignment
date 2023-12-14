package commands

import (
	"testing"

	"github.com/KanhaiyaKumarGupta/inito_Assignment/commands"
	"github.com/stretchr/testify/assert"
)

func TestTouch(t *testing.T) {
	fs := commands.NewFileSystem()
	fs.Mkdir("/test")

	err := fs.Touch("/test/newfile.txt")
	assert.NoError(t, err, "Touch should not produce an error when creating a new file")
	_, exists := fs.GetDir("/test").Files["newfile.txt"]
	assert.True(t, exists, "New file should be created in the directory")

	err = fs.Touch("/test/newfile.txt")
	assert.Error(t, err, "Touch should produce an error when the file already exists")
}

func TestTouchNonExistentDirectory(t *testing.T) {
	fs := commands.NewFileSystem()

	err := fs.Touch("/nonexistent/newfile.txt")
	assert.Error(t, err, "Touch should produce an error when the directory does not exist")
}
