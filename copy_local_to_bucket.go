package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CopyLocalToBucket(source string, bucketURI string) {
	// Trim gs:// from the bucket name
	bucketFullPath := strings.TrimPrefix(bucketURI, "gs://")
	bucketSegments := strings.SplitN(bucketFullPath, "/", 2)
	bucketName := bucketSegments[0]
	bucketPath := bucketSegments[1]
	bucketObject := ""
	if strings.HasSuffix(bucketPath, "/") {
		// e.g. `cp a/sample1.txt gs://sample/b/ => bucketObject: b/sample1.txt
		// e.g. `cp a/sample1.txt gs://sample/ => bucketObject: ./sample1.txt
		bucketObject = filepath.Dir(bucketPath) + "/" + filepath.Base(source)
	} else {
		// e.g. `cp sample1.txt gs://sample/a/sample2.txt => bucketObject: a/sample2.txt
		// e.g. `cp sample1.txt gs://sample/sample2.txt => bucketObject: sample2.txt
		bucketObject = bucketPath
	}

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("Failed to create client:", err)
		return
	}
	defer client.Close()

	// Read local file
	data, err := os.ReadFile(source)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}

	// Get object handle
	obj := client.Bucket(bucketName).Object(bucketObject)

	// Write data to object
	wc := obj.NewWriter(ctx)
	if _, err = wc.Write(data); err != nil {
		fmt.Println("Failed to write to object:", err)
		return
	}
	if err := wc.Close(); err != nil {
		fmt.Println("Failed to close writer:", err)
		return
	}

	fmt.Fprintln(os.Stderr, "gs://"+bucketName+"/"+bucketObject)
}
