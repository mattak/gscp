package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
)

func SyncLocalToBucket(source string, bucketURI string) {
	// Trim gs:// from the bucket name
	bucketName := strings.TrimPrefix(bucketURI, "gs://")

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("Failed to create client:", err)
		return
	}
	defer client.Close()

	err = filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				data, err := ioutil.ReadFile(path)
				if err != nil {
					fmt.Println("Failed to read file:", err)
					return err
				}

				// Get object handle
				obj := client.Bucket(bucketName).Object(path)

				// Write data to object
				wc := obj.NewWriter(ctx)
				if _, err = wc.Write(data); err != nil {
					fmt.Println("Failed to write to object:", err)
					return err
				}
				if err := wc.Close(); err != nil {
					fmt.Println("Failed to close writer:", err)
					return err
				}
			}
			return nil
		})

	if err != nil {
		fmt.Println("Error walking the path", source, ":", err)
		return
	}
}
