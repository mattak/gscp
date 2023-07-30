package main

import (
	"fmt"
	"path"
)

func CommandRemoveObject(bucketUri string) {
	bucketName, bucketPath := SplitBucketURI(bucketUri)

	// client
	ctx, client := CreateClientContext()
	defer client.Close()

	if err := RemoveObject(ctx, client, bucketName, bucketPath); err != nil {
		EprintlnExit(fmt.Sprintf("ERROR: failed to remove object: %v", err))
	}
	fmt.Printf("gs://%s\n", path.Join(bucketName, bucketPath))
}
