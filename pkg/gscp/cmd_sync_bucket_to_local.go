package gscp

import (
	"cloud.google.com/go/storage"
	"fmt"
	"google.golang.org/api/iterator"
	"os"
	"path/filepath"
	"strings"
)

func CommandSyncBucketToLocal(bucketURI string, destination string) error {
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
			fmt.Fprintln(os.Stderr, "ERROR:", err)
			return err
		}

		// filter
		if !strings.HasPrefix(attrs.Name, bucketPath) {
			fmt.Println("skip:\t", attrs.Name)
			continue
		}

		// read
		data, err := ReadObject(ctx, client, bucketName, attrs.Name)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR: failed to read object: ", err)
			return err
		}

		// write
		dstPath := filepath.Join(destination, attrs.Name)
		err = WriteFile(dstPath, data)
		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR: failed to write file: ", err)
			return err
		}

		fmt.Println("copy", "gs://"+filepath.Join(bucketName, attrs.Name), "=>", dstPath)
	}
	return nil
}
