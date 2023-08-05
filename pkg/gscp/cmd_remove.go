package gscp

import (
	"fmt"
	"os"
	"path"
)

func CommandRemoveObject(bucketUri string) error {
	bucketName, bucketPath := SplitBucketURI(bucketUri)

	// client
	ctx, client := CreateClientContext()
	defer client.Close()

	if err := RemoveObject(ctx, client, bucketName, bucketPath); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to remove: gs://"+bucketName+"/"+bucketPath)
		fmt.Fprintln(os.Stderr, "ERROR: failed to remove object: ", err)
		return err
	}
	fmt.Printf("gs://%s\n", path.Join(bucketName, bucketPath))
	return nil
}
