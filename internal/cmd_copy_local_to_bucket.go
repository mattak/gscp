package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CommandCopyLocalToBucket(source string, bucketURI string) error {
	// Trim gs:// from the bucket name
	bucketName, bucketPath := SplitBucketURI(bucketURI)
	bucketObject := ""
	if strings.HasSuffix(bucketPath, "/") {
		// e.g. `cp a/sample1.txt gs://sample/b/ => bucketObject: b/sample1.txt
		// e.g. `cp a/sample1.txt gs://sample/ => bucketObject: ./sample1.txt
		bucketObject = filepath.Join(filepath.Dir(bucketPath), filepath.Base(source))
	} else {
		// e.g. `cp sample1.txt gs://sample/a/sample2.txt => bucketObject: a/sample2.txt
		// e.g. `cp sample1.txt gs://sample/sample2.txt => bucketObject: sample2.txt
		bucketObject = bucketPath
	}

	// client
	ctx, client := CreateClientContext()
	defer client.Close()

	// read
	data, err := ReadFile(source)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: failed to read file: ", source)
		return err
	}

	// write
	err = WriteObject(ctx, client, bucketName, bucketObject, data)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: failed to write object: ", err)
		return err
	}
	fmt.Println(filepath.Join("gs://", bucketName, bucketObject))
	return nil
}
