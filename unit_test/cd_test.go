package commands

import (
	"testing"

	"github.com/KanhaiyaKumarGupta/inito_Assignment/commands"
	"github.com/stretchr/testify/assert"
)

func TestCdToAbsoluteDirectory(t *testing.T) {
	fs := commands.NewFileSystem()
	fs.Mkdir("/testdir")
	err := fs.Cd("/testdir")
	assert.NoError(t, err, "Cd should not produce an error for an existing directory")
	assert.Equal(t, "/testdir", fs.Current.Name, "Current directory should be '/testdir'")
}

func TestCdToParentDirectory(t *testing.T) {
	fs := commands.NewFileSystem()
	fs.Mkdir("/testdir")
	fs.Cd("/testdir")
	err := fs.Cd("..")
	assert.NoError(t, err, "Cd to parent directory should not produce an error")
	assert.Equal(t, "/", fs.Current.Name, "Current directory should be root after moving to parent")
}

func TestCdToRootDirectory(t *testing.T) {
	fs := commands.NewFileSystem()
	fs.Mkdir("/testdir")
	fs.Cd("/testdir")
	err := fs.Cd("/")
	assert.NoError(t, err, "Cd to root directory should not produce an error")
	assert.Equal(t, "/", fs.Current.Name, "Current directory should be root")
}

func TestCdToRootDirectoryWithTilde(t *testing.T) {
	fs := commands.NewFileSystem()
	fs.Mkdir("/testdir")
	fs.Cd("/testdir")
	err := fs.Cd("~")
	assert.NoError(t, err, "Cd to root directory with '~' should not produce an error")
	assert.Equal(t, "/", fs.Current.Name, "Current directory should be root")
}

func TestCdToNonExistentDirectory(t *testing.T) {
	fs := commands.NewFileSystem()
	err := fs.Cd("/nonexistent")
	assert.Error(t, err, "Cd to a non-existent directory should produce an error")
}
