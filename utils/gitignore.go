package utils

import (
	"github.com/denormal/go-gitignore"
	// "log"
	"os"
	// "path/filepath"
	"io/ioutil"
)
type Ignore struct {
	ignore     gitignore.GitIgnore
	ignoreFile string
	delete     bool
}
func InitIgnore(ignoreFile string) (*Ignore, error) {
	ignore, err := gitignore.NewFromFile(ignoreFile)
	if err != nil {
		return nil, err
	}
	return &Ignore{ignore, ignoreFile, false}, nil
}
func InitIgnoreFromText(ignoreText string) (*Ignore, error) {
	err := ioutil.WriteFile("SPACEIgnore", []byte(ignoreText), 0644)
    if err != nil {
		return nil, err
	}
	ignore, err := gitignore.NewFromFile("SPACEIgnore")
	if err != nil {
		return nil, err
	}
	return &Ignore{ignore, "SPACEIgnore", true}, nil
}
func (i *Ignore) ToBeIncluded(path string) bool {
	// Any other checks
	if i.ignore.Ignore(path) {
		return false
	}
	// Any other checks
	return true
}
func (i *Ignore) Close() {
	if(i.delete) {
		os.Remove(i.ignoreFile)
	}
}
// func main() {
// 	i, err := InitIgnoreFromText("")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer i.Close()

// 	var files []string
// 	err = filepath.Walk("test_folder", func(path string, info os.FileInfo, err error) error {
// 		files = append(files, path)
// 		log.Println(path)
// 		log.Println(i.ToBeIncluded(path))
// 		return nil
// 	})
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
