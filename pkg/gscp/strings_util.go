package gscp

import "strings"

func SplitBucketURI(uri string) (string, string) {
	bucketFullPath := strings.TrimPrefix(uri, "gs://")
	bucketSplits := strings.SplitN(bucketFullPath, "/", 2)
	bucketName := bucketSplits[0]

	if len(bucketSplits) >= 2 {
		return bucketName, bucketSplits[1]
	} else {
		return bucketName, ""
	}
}
