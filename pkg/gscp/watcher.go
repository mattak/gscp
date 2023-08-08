package gscp

import (
	"path"
	"time"
)

func UpdateChangedFiles(dirname string, fileMap map[string]int64) ([]string, error) {
	infos, err := ListFiles(dirname)
	if err != nil {
		return nil, err
	}

	changed := []string{}

	for _, info := range infos {
		filePath := path.Join(dirname, info.Name())
		fileTimestamp := info.ModTime().UnixMilli()

		if previousTimestamp, ok := fileMap[filePath]; ok {
			if fileTimestamp > previousTimestamp {
				fileMap[filePath] = fileTimestamp
				changed = append(changed, filePath)
			}
		} else {
			fileMap[filePath] = fileTimestamp
			changed = append(changed, filePath)
		}
	}

	return changed, nil
}

func WatchDir(waitTime time.Duration, dirname string, onUpdate func(files []string)) error {
	fileMap := map[string]int64{}

	for {
		time.Sleep(waitTime)
		changedFiles, err := UpdateChangedFiles(dirname, fileMap)

		if err != nil {
			return err
		}
		if len(changedFiles) < 1 {
			continue
		}

		onUpdate(changedFiles)
	}
}
