package gscp

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func ReadObject(
	ctx *context.Context,
	client *storage.Client,
	bucketName string,
	objectName string,
) ([]byte, error) {
	rc, err := client.Bucket(bucketName).Object(objectName).NewReader(*ctx)
	defer rc.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to open object: ", err)
		return nil, err
	}

	data, err := ioutil.ReadAll(rc)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to read object: ", err)
		return nil, err
	}

	return data, nil
}

func WriteObject(
	ctx *context.Context,
	client *storage.Client,
	bucketName string,
	objectPath string,
	data []byte,
) error {
	// Get object handle
	obj := client.Bucket(bucketName).Object(objectPath)

	// Write data to object
	wc := obj.NewWriter(*ctx)
	if _, err := wc.Write(data); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to write to object: ", err)
		return err
	}
	if err := wc.Close(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to close writer: ", err)
		return err
	}

	return nil
}

func RemoveObject(
	ctx *context.Context,
	client *storage.Client,
	bucketName string,
	objectPath string,
) error {
	obj := client.Bucket(bucketName).Object(objectPath)
	if err := obj.Delete(*ctx); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to delete object: ", err)
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

type DetailObject struct {
	Path    string
	Created time.Time
	Updated time.Time
	Size    int64
}

func ListDetailObject(
	ctx *context.Context,
	client *storage.Client,
	bucketName string,
	objectPath string,
) ([]DetailObject, error) {
	it := client.Bucket(bucketName).Objects(*ctx, nil)
	results := []DetailObject{}

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		if strings.HasPrefix(attrs.Name, objectPath) {
			o := DetailObject{
				Path:    attrs.Name,
				Created: attrs.Created,
				Updated: attrs.Created,
				Size:    attrs.Size,
			}
			results = append(results, o)
		}
	}

	return results, nil
}

func CreateClientContext() (*context.Context, *storage.Client) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to create client:", err)
		return nil, nil
	}
	return &ctx, client
}
