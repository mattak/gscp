package main

import (
	"fmt"
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

func ReadFile(source string) ([]byte, error) {
	data, err := os.ReadFile(source)
	if err != nil {
		fmt.Errorf("Failed to read file: %v\n", err)
		return nil, err
	}
	return data, nil
}

func WriteFile(path string, data []byte) error {
	dir := filepath.Dir(path)
	if !IsDirExist(dir) {
		if err := mkdirp(dir); err != nil {
			fmt.Errorf("Failed to create directory: %s\n", dir)
			return err
		}
	}

	err := os.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Errorf("Failed to write file: %v\n", err)
		return err
	}

	return nil
}
