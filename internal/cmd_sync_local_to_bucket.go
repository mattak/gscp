package internal

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func CommandSyncLocalToBucket(source string, bucketURI string) error {
	bucketName, bucketPath := SplitBucketURI(bucketURI)

	// client
	ctx, client := CreateClientContext()
	defer client.Close()

	err := filepath.WalkDir(source,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}

			// read
			data, err := ReadFile(path)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERROR: failed to read file: ", err)
				return err
			}

			// write
			pathForObject := strings.TrimPrefix(path, source)
			objectPath := filepath.Join(bucketPath, pathForObject)
			err = WriteObject(ctx, client, bucketName, objectPath, data)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ERROR: failed to write object: ", err)
				return err
			}
			fmt.Println("copy", path, "=>", "gs://"+filepath.Join(bucketName, objectPath))
			return err
		})

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: walking the path", source, ":", err)
		return err
	}

	return nil
}
