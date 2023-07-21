package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"strings"
)

// list all objects in a bucket
func ListObjects(bucketURI string) {
	// Trim gs:// from the bucket name

	bucketFullPath := strings.TrimPrefix(bucketURI, "gs://")
	bucketSegments := strings.SplitN(bucketFullPath, "/", 2)
	bucketName := bucketSegments[0]
	bucketPath := bucketSegments[1]

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
		if strings.HasPrefix(attrs.Name, bucketPath) {
			fmt.Println(attrs.Name)
		}
	}
}
