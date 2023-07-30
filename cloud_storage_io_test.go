package main

import (
	"os"
	"path/filepath"
	"testing"
)

func SetupBucketEnv() (string, string) {
	ENV_TEST_BUCKET_NAME := os.Getenv("ENV_TEST_BUCKET_NAME")
	ENV_TEST_BUCKET_PATH := os.Getenv("ENV_TEST_BUCKET_PATH")
	return ENV_TEST_BUCKET_NAME, ENV_TEST_BUCKET_PATH
}

func setup_cloud_storage_io() {
}

func teardown_cloud_storage_io() {
}

func TestWriteReadObject(t *testing.T) {
	setup_cloud_storage_io()
	bucketName, bucketPath := SetupBucketEnv()

	tf := TestFrame{Test: t}

	tf.Run("write & read & remove object", func(tf TestFrame) {
		ctx, client := CreateClientContext()
		defer client.Close()

		path := filepath.Join(bucketPath, "a.txt")
		err := WriteObject(ctx, client, bucketName, path, []byte("hello"))
		tf.AssertNil(err)

		data, err := ReadObject(ctx, client, bucketName, path)
		tf.AssertNil(err)
		tf.AssertEquals(string(data), "hello")

		err = RemoveObject(ctx, client, bucketName, path)
		tf.AssertNil(err)
	})

	teardown_cloud_storage_io()
}
