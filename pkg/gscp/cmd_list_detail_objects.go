package gscp

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type CommandListDetailObjectsOption struct {
	WithLatest1    bool
	WithNameOnly   bool
	WithUnixMillis bool
}

func getTimeFormat(t time.Time, option CommandListDetailObjectsOption) string {
	if option.WithUnixMillis {
		return strconv.FormatInt(t.UnixMilli(), 10)
	}
	return t.Format("2006-01-02_15:04:05-0700")
}

func printDetailObject(o DetailObject, option CommandListDetailObjectsOption) {
	if option.WithNameOnly {
		fmt.Println(o.Path)
	} else {
		line := strings.Join(
			[]string{
				getTimeFormat(o.Created, option),
				getTimeFormat(o.Updated, option),
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
func CommandListDetailObjects(bucketURI string, option CommandListDetailObjectsOption) error {
	bucketName, bucketPath := SplitBucketURI(bucketURI)

	// client
	ctx, client := CreateClientContext()
	defer client.Close()

	objects, err := ListDetailObject(ctx, client, bucketName, bucketPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR: list object failed: ", err)
		return err
	}

	if option.WithLatest1 {
		if len(objects) < 1 {
			return nil
		}
		maxIndex := findLatestDetailObjectIndex(objects)
		if maxIndex < 0 {
			return nil
		}

		printDetailObject(objects[maxIndex], option)
	} else {
		for _, o := range objects {
			printDetailObject(o, option)
		}
	}

	return nil
}
