package main

import (
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

func RemoveObject(objectURI string) {
	parts := strings.SplitN(strings.TrimPrefix(objectURI, "gs://"), "/", 2)
	bucketName := parts[0]
	objectName := parts[1]

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("Failed to create client:", err)
		os.Exit(1)
	}
	defer client.Close()

	// Delete the object
	if err := client.Bucket(bucketName).Object(objectName).Delete(ctx); err != nil {
		fmt.Println("Failed to delete object:", err)
		os.Exit(1)
	}

	fmt.Println("gs://" + bucketName + "/" + objectName)
}
