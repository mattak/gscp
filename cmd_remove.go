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
		Eprintln("Failed to remove: gs://" + bucketName + "/" + bucketPath)
		EprintlnExit("ERROR: failed to remove object: ", err)
		return
	}
	fmt.Printf("gs://%s\n", path.Join(bucketName, bucketPath))
}
