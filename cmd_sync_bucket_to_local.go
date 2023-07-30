package main

import (
	"cloud.google.com/go/storage"
	"fmt"
	"path/filepath"
	"strings"

	"google.golang.org/api/iterator"
)

func CommandSyncBucketToLocal(bucketURI string, destination string) {
	bucketName, bucketPath := SplitBucketURI(bucketURI)

	// client
	ctx, client := CreateClientContext()
	defer client.Close()

	q := storage.Query{Prefix: bucketPath}
	it := client.Bucket(bucketName).Objects(*ctx, &q)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			EprintlnExit("Error:", err)
			return
		}

		// filter
		if !strings.HasPrefix(attrs.Name, bucketPath) {
			fmt.Println("skip:\t", attrs.Name)
			continue
		}

		// read
		data := ReadObject(ctx, client, bucketName, attrs.Name)

		// write
		dstPath := filepath.Join(destination, attrs.Name)
		WriteFile(dstPath, data)
		fmt.Println("write", filepath.Join("gs://", bucketName, attrs.Name), "=>", dstPath)
	}
}
