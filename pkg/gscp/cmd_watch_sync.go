package gscp

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"time"
)

func CopyObject(
	ctx *context.Context,
	client *storage.Client,
	sourceFilepath string,
	bucketName, bucketObject string,
) error {
	// read
	data, err := ReadFile(sourceFilepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: failed to read file: ", sourceFilepath)
		return err
	}

	// write
	err = WriteObject(ctx, client, bucketName, bucketObject, data)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: failed to write object: ", err)
		return err
	}

	return nil
}

func watchToCopyObject(
	ctx *context.Context,
	client *storage.Client,
	sourceDir string,
	bucketName string,
	bucketPath string,
) error {

	err := WatchDir(1*time.Second, sourceDir, func(files []string) {
		for _, file := range files {
			objectName := path.Base(file)
			objectPath := path.Join(bucketPath, objectName)

			bucketUri := strings.Join([]string{"gs:/", bucketName, objectPath}, "/")
			err := CopyObject(ctx, client, file, bucketName, objectPath)
			if err != nil {
				log.Fatalln("ERROR: CopyObject: ", file, " => ", bucketUri, err)
			}
			fmt.Println("copy", file, " => ", bucketUri)
		}
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: Watch dir", sourceDir, ":", err)
		return err
	}

	return nil
}

func CommandWatchSync(sourceDir string, bucketURI string) error {
	bucketName, bucketPath := SplitBucketURI(bucketURI)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		// client
		ctx, client := CreateClientContext()
		defer client.Close()

		for {
			err := watchToCopyObject(ctx, client, sourceDir, bucketName, bucketPath)
			if err != nil {
				log.Fatalln("ERROR: watch loop: ", err)
			}
		}
	}()

	sig := <-ch
	fmt.Println("Exit to catch signal: ", sig)
	return nil
}
