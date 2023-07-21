package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

func SyncBucketToLocal(bucketURI string, destination string) {
	// Trim gs:// from the bucket name
	bucketName := strings.TrimPrefix(bucketURI, "gs://")

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("Failed to create client:", err)
		return
	}
	defer client.Close()

	it := client.Bucket(bucketName).Objects(ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Open the object
		rc, err := client.Bucket(bucketName).Object(attrs.Name).NewReader(ctx)
		if err != nil {
			fmt.Println("Failed to open object:", err)
			return
		}

		data, err := ioutil.ReadAll(rc)
		if err != nil {
			fmt.Println("Failed to read object:", err)
			return
		}

		rc.Close()

		// Create the file
		err = os.WriteFile(filepath.Join(destination, attrs.Name), data, 0644)
		if err != nil {
			fmt.Println("Failed to write file:", err)
			return
		}
	}
}
