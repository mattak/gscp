package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// copy object to/from bucket
func CommandCopyObjectToLocal(bucketURI string, destination string) {
	bucketName, objectName := SplitBucketURI(bucketURI)

	if IsDir(destination) {
		destination = filepath.Join(filepath.Dir(destination), filepath.Base(objectName))
	}

	//  client
	ctx, client := CreateClientContext()
	defer client.Close()

	// read
	data, err := ReadObject(ctx, client, bucketName, objectName)
	if err != nil {
		EprintlnExit("ERROR: failed to read object: ", err)
		return
	}

	// write
	err = WriteFile(destination, data)
	if err != nil {
		EprintlnExit("ERROR: failed to write file: ", err)
	}
	fmt.Fprintln(os.Stderr, destination)
}
