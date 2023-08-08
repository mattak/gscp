package gscp

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

func IsDir(path string) bool {
	dir := filepath.Base(path)
	if dir == "." || dir == ".." || strings.HasSuffix(path, "/") {
		return true
	}
	return false
}

func mkdirp(path string) error {
	return os.MkdirAll(path, 0755)
}

func ReadFile(source string) ([]byte, error) {
	data, err := os.ReadFile(source)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read file: ", err)
		return nil, err
	}
	return data, nil
}

func WriteFile(path string, data []byte) error {
	dir := filepath.Dir(path)
	if !IsDirExist(dir) {
		if err := mkdirp(dir); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create directory: ", dir)
			return err
		}
	}

	err := os.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write file: ", err)
		return err
	}

	return nil
}

func ListFiles(dirname string) ([]fs.FileInfo, error) {
	entries, err := os.ReadDir(dirname)
	if err != nil {
		return nil, err
	}

	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		infos = append(infos, info)
	}
	return infos, nil
}
