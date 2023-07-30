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
			EprintlnExit("ERROR:", err)
			return
		}

		// filter
		if !strings.HasPrefix(attrs.Name, bucketPath) {
			fmt.Println("skip:\t", attrs.Name)
			continue
		}

		// read
		data, err := ReadObject(ctx, client, bucketName, attrs.Name)
		if err != nil {
			EprintlnExit("ERROR: failed to read object: ", err)
			return
		}

		// write
		dstPath := filepath.Join(destination, attrs.Name)
		WriteFile(dstPath, data)
		fmt.Println("copy", "gs://"+filepath.Join(bucketName, attrs.Name), "=>", dstPath)
	}
}
