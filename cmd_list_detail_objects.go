package main

import (
	"fmt"
	"strconv"
	"strings"
)

type CommandListDetailObjectsOption struct {
	WithLatest1  bool
	WithNameOnly bool
}

func printDetailObject(o DetailObject, option CommandListDetailObjectsOption) {
	if option.WithNameOnly {
		fmt.Println(o.Path)
	} else {
		line := strings.Join(
			[]string{
				o.Created.Format("2006-01-02_15:04:05-0700"),
				o.Updated.Format("2006-01-02_15:04:05-0700"),
				strconv.FormatInt(o.Size, 10),
				o.Path,
			},
			"\t",
		)
		fmt.Println(line)
	}
}

func findLatestDetailObjectIndex(objects []DetailObject) int {
	maxValue := int64(0)
	maxIndex := -1
	for i, v := range objects {
		t := v.Updated.UnixMilli()
		if maxValue <= t {
			maxValue = t
			maxIndex = i
		}
	}
	return maxIndex
}

// list all detailed objects in a bucket
func CommandListDetailObjects(bucketURI string, option CommandListDetailObjectsOption) {
	bucketName, bucketPath := SplitBucketURI(bucketURI)

	// client
	ctx, client := CreateClientContext()
	defer client.Close()

	objects, err := ListDetailObject(ctx, client, bucketName, bucketPath)
	if err != nil {
		EprintlnExit("ERROR: list object failed: ", err)
		return
	}

	if option.WithLatest1 {
		if len(objects) < 1 {
			return
		}
		maxIndex := findLatestDetailObjectIndex(objects)
		if maxIndex < 0 {
			return
		}

		printDetailObject(objects[maxIndex], option)
	} else {
		for _, o := range objects {
			printDetailObject(o, option)
		}
	}
}
