package gscp

import (
	"fmt"
	"os"
)

// list all objects in a bucket
func CommandListObjects(bucketURI string) error {
	bucketName, bucketPath := SplitBucketURI(bucketURI)

	// client
	ctx, client := CreateClientContext()
	defer client.Close()

	paths, err := ListObject(ctx, client, bucketName, bucketPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: list detail objects failed: ", err)
		return err
	}

	for _, p := range paths {
		fmt.Println(p)
	}

	return nil
}
