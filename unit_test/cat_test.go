package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCat(t *testing.T) {
	fs := setupFileSystemWithFiles()

	content, err := fs.Cat("/test/file1.txt")
	assert.NoError(t, err, "Cat should not produce an error for an existing file")
	assert.Equal(t, "Hello World", content, "Content of 'file1.txt' should match")

	_, err = fs.Cat("/test/nonexistent.txt")
	assert.Error(t, err, "Cat should produce an error for a non-existent file")

	_, err = fs.Cat("/nonexistent/file.txt")
	assert.Error(t, err, "Cat should produce an error when the directory does not exist")
}
