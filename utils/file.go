package utils

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func SaveFileAndGetMd5(src io.Reader, fileDir string) (string, string, error) {
	filePath := filepath.Join(fileDir, Base62UUID())
	dst, err := os.Create(filePath)
	if err != nil {
		return "", "", err
	}
	defer dst.Close()

	hashMd5 := md5.New()
	tr := io.TeeReader(src, hashMd5)
	_, err = io.Copy(dst, tr)
	if err != nil {
		return "", "", err
	}
	return filePath, hex.EncodeToString(hashMd5.Sum(nil)), nil
}

func SaveFileWithMd5Name(src io.Reader, fileDir string) (string, error) {
	filePath, fileMd5, err := SaveFileAndGetMd5(src, fileDir)
	if err != nil {
		return "", err
	}
	md5FilePath := filepath.Join(fileDir, fileMd5)
	return fileMd5, os.Rename(filePath, md5FilePath)
}

func GetFileMd5(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	h := md5.New()

	_, err = io.Copy(h, r)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

func RenameFileToMd5(filePath string) (string, error) {
	fileMd5, err := GetFileMd5(filePath)
	if err != nil {
		return "", err
	}
	fileDir := filepath.Dir(filePath)
	newFilePath := filepath.Join(fileDir, fileMd5)
	err = os.Rename(filePath, newFilePath)
	if err != nil {
		return "", err
	}
	return newFilePath, nil
}

func ListFilePathRecursive(root string) ([]string, error) {
	var filePaths []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	return filePaths, err
}

func ListFileName(root string) ([]string, error) {
	files, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames, nil
}

func GetFileModTime(filePath string) (time.Time, error) {
	fileStat, err := os.Stat(filePath)
	if err != nil {
		return time.Time{}, err
	}
	return fileStat.ModTime(), nil
}

// check by ext
func IsImgFile(filePath string) bool {
	imgExts := map[string]bool{
		"jpg":  true,
		"jpeg": true,
		"png":  true,
		"bmp":  true,
	}
	return imgExts[strings.ToLower(filepath.Ext(filePath))]
}
