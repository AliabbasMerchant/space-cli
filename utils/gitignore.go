package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/denormal/go-gitignore"
)

// Ignore is the instance used to selectively add files to archive
type Ignore struct {
	ignore     gitignore.GitIgnore
	ignoreFile string
	delete     bool
}

// NewIgnore creates a new Ignore instance
func NewIgnore(ignoreFile string) (*Ignore, error) {
	ignore, err := gitignore.NewFromFile(ignoreFile)
	if err != nil {
		return nil, err
	}
	return &Ignore{ignore, ignoreFile, false}, nil
}

// GetFileList fetches the file list after applying the ignore rule
func (i *Ignore) GetFileList(workingDir string) ([]string, error) {
	// Change the directory to the working directory
	if err := os.Chdir(workingDir); err != nil {
		return nil, err
	}

	// Create a files object
	files := []string{}

	// Walk through the file tree
	if err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		// Add the file to list if it passes the match
		if !info.IsDir() && i.include(path) {
			log.Println("Adding file:", path)
			files = append(files, path)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}

func (i *Ignore) include(path string) bool {
	return !i.ignore.Ignore(path)
}
