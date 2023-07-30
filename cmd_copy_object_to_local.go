package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// copy object to/from bucket
func CommandCopyObjectToLocal(bucketURI string, destination string) {
	bucketName, objectName := SplitBucketURI(bucketURI)

	if destination == "." {
		destination = filepath.Base(objectName)
	}

	//  client
	ctx, client := CreateClientContext()
	defer client.Close()

	// read
	data := ReadObject(ctx, client, bucketName, objectName)

	// write
	WriteFile(destination, data)
	fmt.Fprintln(os.Stderr, destination)
}
