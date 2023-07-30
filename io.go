package main

import (
	"os"
	"path/filepath"
)

func IsDirExist(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	} else if err != nil {
		return false
	} else if info.IsDir() {
		return true
	} else {
		return false
	}
}

func mkdirp(path string) error {
	return os.MkdirAll(path, 0755)
}

func ReadFile(source string) []byte {
	data, err := os.ReadFile(source)
	if err != nil {
		EprintlnExit("Failed to read file:", err)
		return nil
	}
	return data
}

func WriteFile(path string, data []byte) {
	dir := filepath.Dir(path)
	if !IsDirExist(dir) {
		if err := mkdirp(dir); err != nil {
			EprintlnExit("Failed to create directory: " + dir)
		}
	}

	err := os.WriteFile(path, data, 0644)
	if err != nil {
		EprintlnExit("Failed to write file:", err)
		return
	}
}
