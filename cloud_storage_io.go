package main

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"strings"
)

func ReadObject(ctx *context.Context, client *storage.Client, bucketName string, objectName string) []byte {
	rc, err := client.Bucket(bucketName).Object(objectName).NewReader(*ctx)
	defer rc.Close()

	if err != nil {
		EprintlnExit("Failed to open object:", err)
		return nil
	}

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		EprintlnExit("Failed to read object:", err)
		return nil
	}

	return data
}

func WriteObject(ctx *context.Context, client *storage.Client, bucketName string, objectPath string, data []byte) {
	// Get object handle
	obj := client.Bucket(bucketName).Object(objectPath)

	// Write data to object
	wc := obj.NewWriter(*ctx)
	if _, err := wc.Write(data); err != nil {
		EprintlnExit("Failed to write to object:", err)
	}
	if err := wc.Close(); err != nil {
		EprintlnExit("Failed to close writer:", err)
	}
}

func RemoveObject(
	ctx *context.Context,
	client *storage.Client,
	bucketName string,
	objectPath string,
) error {
	obj := client.Bucket(bucketName).Object(objectPath)
	if err := obj.Delete(*ctx); err != nil {
		fmt.Errorf("Failed to delete object: %v\n", err)
		return err
	}
	return nil
}

func ListObject(
	ctx *context.Context,
	client *storage.Client,
	bucketName string,
	objectPath string,
) ([]string, error) {
	it := client.Bucket(bucketName).Objects(*ctx, nil)
	results := []string{}

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		if strings.HasPrefix(attrs.Name, objectPath) {
			results = append(results, attrs.Name)
		}
	}

	return results, nil
}

func CreateClientContext() (*context.Context, *storage.Client) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		EprintlnExit("Failed to create client:", err)
		return nil, nil
	}
	return &ctx, client
}
