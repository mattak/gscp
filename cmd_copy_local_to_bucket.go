package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func CommandCopyLocalToBucket(source string, bucketURI string) {
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
	data := ReadFile(source)

	// write
	WriteObject(ctx, client, bucketName, bucketObject, data)
	fmt.Println(filepath.Join("gs://", bucketName, bucketObject))
}
