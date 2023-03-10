package utils

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Decompress a .tar.gz file into a directory
func DecompressTarGz(srcTarGzFile, dstDirPath string) error {
	// Open the tar.gz file
	file, err := os.Open(srcTarGzFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a gzip reader for the file
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	// Create a tar reader for the gzip reader
	tarReader := tar.NewReader(gzipReader)

	// Loop through each file in the tar archive
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			// End of tar archive
			break
		}
		if err != nil {
			return err
		}

		targetFilePath := filepath.Join(dstDirPath, header.Name)

		// Create sub directories if needed
		info := header.FileInfo()
		if err = os.MkdirAll(filepath.Dir(targetFilePath), info.Mode()); err != nil {
			return err
		}

		// Create a new file for the current file in the archive
		file, err := os.OpenFile(targetFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return err
		}
		defer file.Close()

		// Write the contents of the current file to the new file
		if _, err := io.Copy(file, tarReader); err != nil {
			return err
		}
	}

	return nil
}

// compress a directory into a .tar.gz file
func CreateTarGz(srcDirPath, dstTarGzFilePath string) error {
	// Normalize the path format
	srcDirPath = filepath.Clean(srcDirPath)
	dstTarGzFilePath = filepath.Clean(dstTarGzFilePath)

	// Create the tar.gz file
	file, err := os.Create(dstTarGzFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a gzip writer for the file
	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	// Create a tar writer for the gzip writer
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// Compress the directory
	err = filepath.Walk(srcDirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create a new tar header for the file
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		// Fix file name inside the tar.gz
		header.Name = strings.TrimPrefix(strings.TrimPrefix(path, srcDirPath), "\\")

		// Ignore the most outside directory
		if header.Name == "" {
			return nil
		}

		// Write the header to the tar archive
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// If the file is not a directory, write the contents of the file to the tar archive
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			if _, err := io.Copy(tarWriter, file); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
