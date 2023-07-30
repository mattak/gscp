package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

func CommandSyncLocalToBucket(source string, bucketURI string) {
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
			data := ReadFile(path)

			// write
			pathForObject := strings.TrimPrefix(path, source)
			objectPath := filepath.Join(bucketPath, pathForObject)
			WriteObject(ctx, client, bucketName, objectPath, data)
			fmt.Println("write:", path, "=>", filepath.Join("gs://", bucketName, objectPath))
			return err
		})

	if err != nil {
		EprintlnExit("Error walking the path", source, ":", err)
		return
	}
}
