package utils

import (
	"archive/zip"
	"io"
	"os"
)

// ZipFiles creates an archive based on list of file paths provided
func ZipFiles(filename string, files []string) error {
	// Create a zip archive
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()

	// Create a zip writer
	zipWriter := zip.NewWriter(newZipFile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err := addFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	return nil
}

func addFileToZip(zipWriter *zip.Writer, filename string) error {
	// Open the file
	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get file stats
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	// Set the file name in header
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}
	header.Name = filename

	// Set compression method
	header.Method = zip.Deflate

	// Create a writer
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}

	// Copy the file to archive
	_, err = io.Copy(writer, fileToZip)
	return err
}

// func main() {
//   err := ZipFiles("a.zip", []string{"test_folder\\.ignore", "test_folder\\a"})
//   fmt.Println(err)
// }
