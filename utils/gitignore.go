package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/monochromegane/go-gitignore"
)

// Ignore is the instance used to selectively add files to archive
type Ignore struct {
	ignore     gitignore.IgnoreMatcher
	ignoreFile string
	delete     bool
}

// NewIgnore creates a new Ignore instance
func NewIgnore(ignoreFile string) (*Ignore, error) {
	ignore, err := gitignore.NewGitIgnore(ignoreFile)
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
	ignored := []string{".git"}

	// Walk through the file tree
	if err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		// Add the file to list if it passes the match
		if i.ignore.Match(path, info.IsDir()) {
			log.Println("Ignoring file:", path)
			ignored = append(ignored, path)
		}
		if !info.IsDir() && !contains(ignored, path) {
			log.Println("Adding file:", path)
			files = append(files, path)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.HasPrefix(e, a) {
			return true
		}
	}
	return false
}
