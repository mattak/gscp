package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// copy object to/from bucket
func CopyObjectToLocal(bucketURI string, destination string) {
	// Trim gs:// from the bucket name
	bucketName := strings.TrimPrefix(bucketURI, "gs://")
	objectNames := strings.SplitN(bucketName, "/", 2)
	if len(objectNames) != 2 {
		fmt.Println("Failed to parse bucket and object name")
		return
	}
	bucketName = objectNames[0]
	objectName := objectNames[1]

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("Failed to create client:", err)
		return
	}
	defer client.Close()

	rc, err := client.Bucket(bucketName).Object(objectName).NewReader(ctx)
	if err != nil {
		fmt.Println("Failed to open object:", err)
		return
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		fmt.Println("Failed to read object:", err)
		return
	}

	if destination == "." {
		destination = filepath.Base(objectName)
	}

	err = os.WriteFile(destination, data, 0644)
	if err != nil {
		fmt.Println("Failed to write file:", destination, err)
		return
	}

	fmt.Fprintln(os.Stderr, destination)
}
