package utils

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// suit for run by go build or go run
func GetCurrentAbPath() string {
	curDir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(curDir, tmpDir) {
		return getCurrentAbPathByCaller()
	}
	return curDir
}

// suit for run by go build
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}
	abPath, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return abPath
}

// suit for run by go run
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func IsFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func GetTmpFilename() (filename string) {
	return "/tmp/" + Base62UUID()
}
