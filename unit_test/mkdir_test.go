package commands

import (
	"testing"

	"github.com/KanhaiyaKumarGupta/inito_Assignment/commands"
	"github.com/stretchr/testify/assert"
)

func TestMkdirSingleDirectory(t *testing.T) {
	fs := commands.NewFileSystem()

	err := fs.Mkdir("/newdir")
	assert.NoError(t, err, "Mkdir should not produce an error for a new directory")

	newDir := fs.GetDir("/newdir")
	assert.NotNil(t, newDir, "Directory '/newdir' should be created")
}

func TestMkdirNestedDirectories(t *testing.T) {
	fs := commands.NewFileSystem()

	err := fs.Mkdir("/nested/dir/structure")
	assert.NoError(t, err, "Mkdir should not produce an error for nested directories")

	nestedDir := fs.GetDir("/nested/dir/structure")
	assert.NotNil(t, nestedDir, "Nested directories '/nested/dir/structure' should be created")
}

func TestMkdirExistingDirectory(t *testing.T) {
	fs := commands.NewFileSystem()
	fs.Mkdir("/existing")

	err := fs.Mkdir("/existing")

	assert.Error(t, err, "Mkdir should produce an error when creating an existing directory")
}
