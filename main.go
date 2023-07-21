package main

import (
	"fmt"
	"os"
	"strings"
)

func eprintlnExit(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		eprintlnExit("Error: requires command name")
		return
	}

	command := os.Args[1]
	switch command {
	case "ls":
		CheckEnvironmentValues()
		uri := os.Args[2]
		ListObjects(uri)
	case "cp":
		CheckEnvironmentValues()
		if len(os.Args) < 4 {
			eprintlnExit("Error: cp requires 2 arguments")
		}

		srcUri := os.Args[2]
		dstUri := os.Args[3]
		if strings.HasPrefix(srcUri, "gs://") {
			CopyObjectToLocal(srcUri, dstUri)
		} else {
			CopyLocalToBucket(srcUri, dstUri)
		}
	case "rm":
		CheckEnvironmentValues()
		uri := os.Args[2]
		RemoveObject(uri)
	case "rsync":
		CheckEnvironmentValues()

		srcUri := os.Args[2]
		dstUri := os.Args[3]
		if strings.HasPrefix(srcUri, "gs://") {
			SyncBucketToLocal(srcUri, dstUri)
		} else {
			SyncLocalToBucket(srcUri, dstUri)
		}
	default:
		fmt.Println("Invalid command")
	}
}
