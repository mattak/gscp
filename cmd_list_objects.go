package main

import "fmt"

// list all objects in a bucket
func CommandListObjects(bucketURI string) {
	bucketName, bucketPath := SplitBucketURI(bucketURI)

	// client
	ctx, client := CreateClientContext()
	defer client.Close()

	paths, err := ListObject(ctx, client, bucketName, bucketPath)
	if err != nil {
		EprintlnExit("ERROR: list object failed: ", err)
		return
	}

	for _, p := range paths {
		fmt.Println(p)
	}
}
