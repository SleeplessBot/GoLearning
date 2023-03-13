// content addressable storage
package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const casDir = ".cas"

type Cas interface {
	BaseDir() string
	IsExist(key string) bool
	Create(io.Reader) (key string, err error)
	Delete(key string) (err error)
	GetFilePath(key string) (filePath string, err error)
	MoveFileToCas(filePath string) (newFilePath string, key string, err error)
	FilePathToKey(filePath string) (key string, err error)
}

var DefaultCas Cas = SimpleCas{}

type SimpleCas struct{}

func (s SimpleCas) BaseDir() string {
	return casDir
}

func (s SimpleCas) IsExist(key string) bool {
	filePath, err := s.GetFilePath(key)
	if err != nil {
		return false
	}
	return FileExists(filePath)
}

func (s SimpleCas) Create(file io.Reader) (key string, err error) {
	return SaveFileWithSha1Name(file, s.BaseDir())
}

func (s SimpleCas) Delete(key string) (err error) {
	filePath, err := s.GetFilePath(key)
	if err != nil {
		return err
	}
	return os.Remove(filePath)
}

func (s SimpleCas) GetFilePath(key string) (filePath string, err error) {
	key = strings.ToLower(key)
	if !IsSha1LikeHashString(key) {
		return "", fmt.Errorf("invalid key")
	}
	return filepath.Join(s.BaseDir(), key[:2], key[2:]), nil
}

func (s SimpleCas) MoveFileToCas(filePath string) (newFilePath string, key string, err error) {
	// if file already in cas, do nothing
	key, err = s.FilePathToKey(filePath)
	if err == nil {
		return filePath, key, nil
	}

	key, err = GetFileSha1(filePath)
	if err != nil {
		return "", "", err
	}
	newFilePath = filepath.Join(s.BaseDir(), key[:2], key[2:])
	err = os.MkdirAll(filepath.Dir(newFilePath), os.ModePerm)
	if err != nil {
		return "", "", err
	}
	err = os.Rename(filePath, newFilePath)
	if err != nil {
		return "", "", err
	}
	return newFilePath, key, nil
}

func (s SimpleCas) FilePathToKey(filePath string) (key string, err error) {
	filePath = filepath.Clean(strings.ToLower(filePath))
	pathElements := strings.Split(filePath, string(os.PathSeparator))
	pathElementsLen := len(pathElements)
	if pathElementsLen < 3 {
		return "", fmt.Errorf("invalid file path")
	}
	absFilePath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}
	targetAbsFilePath, err := filepath.Abs(filepath.Join(s.BaseDir(), pathElements[pathElementsLen-2], pathElements[pathElementsLen-1]))
	if err != nil {
		return "", err
	}

	if filepath.Clean(absFilePath) != filepath.Clean(targetAbsFilePath) {
		return "", fmt.Errorf("invalid file path")
	}
	key = pathElements[pathElementsLen-2] + pathElements[pathElementsLen-1]
	if !IsSha1LikeHashString(key) {
		return "", fmt.Errorf("invalid file path")
	}
	return key, nil
}
