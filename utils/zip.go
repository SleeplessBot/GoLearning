package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Decompress a zip file into a directory
func Unzip(srcZipFilePath, dstDirPath string) error {
	// Open the zip file
	zipReader, err := zip.OpenReader(srcZipFilePath)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	// Iterate through the files in the archive
	for _, file := range zipReader.File {
		// Create directory
		fullPath := filepath.Join(dstDirPath, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(fullPath, os.ModePerm)
		} else {
			os.MkdirAll(filepath.Dir(fullPath), os.ModePerm)

			// Create the destination file
			dstFile, err := os.Create(fullPath)
			if err != nil {
				return err
			}
			defer dstFile.Close()

			// Open the file in the archive
			srcFile, err := file.Open()
			if err != nil {
				return err
			}
			defer srcFile.Close()

			// Copy the contents of the file to the destination file
			_, err = io.Copy(dstFile, srcFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// compress a directory into a zip file
func CreateZip(srcDirPath, dstZipFilePath string) error {
	// Create a new zip file
	zipFile, err := os.Create(dstZipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Create a new zip archive
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through all files in the directory and add them to the zip archive
	filepath.Walk(srcDirPath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Dir does not carry any contents, ignore it
		if info.IsDir() {
			return nil
		}

		// Create a new file header
		fileHeader, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Set the name of the file header to the relative path of the file
		fileHeader.Name, err = filepath.Rel(srcDirPath, filePath)
		if err != nil {
			return err
		}

		// Add the file header to the zip archive
		fileWriter, err := zipWriter.CreateHeader(fileHeader)
		if err != nil {
			return err
		}

		// Add the file contents to the zip archive
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(fileWriter, file)
		if err != nil {
			return err
		}

		return nil
	})
	return nil
}
