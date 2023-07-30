package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

func EprintlnExit(messages ...any) {
	fmt.Fprintln(os.Stderr, messages...)
	os.Exit(1)
}

func main() {
	app := cli.App{
		Name:  "gscp",
		Usage: "minimum set of google cloud copy command",
		Commands: []*cli.Command{
			{
				Name:        "ls",
				Description: "list up files of google cloud storage.",
				ArgsUsage:   "[bucketUri:\"gs://bucketName/path\"]",
				Action: func(ctx *cli.Context) error {
					CheckEnvironmentValues()

					if ctx.Args().Len() < 1 {
						EprintlnExit("ERROR: ls requires a argument")
						return nil
					}

					uri := ctx.Args().First()
					CommandListObjects(uri)
					return nil
				},
			},
			{
				Name:        "cp",
				Description: "copy files between google cloud storage files [gs://bucketName/path] and local files [/path/to/file].",
				ArgsUsage:   "[bucketUri|localPath] [localPath|bucketUri]",
				Action: func(ctx *cli.Context) error {
					CheckEnvironmentValues()

					if ctx.Args().Len() < 2 {
						EprintlnExit("ERROR: cp requires 2 arguments")
						return nil
					}

					srcUri := ctx.Args().Get(0)
					dstUri := ctx.Args().Get(1)
					if strings.HasPrefix(srcUri, "gs://") {
						CommandCopyObjectToLocal(srcUri, dstUri)
					} else {
						CommandCopyLocalToBucket(srcUri, dstUri)
					}
					return nil
				},
			},
			{
				Name:        "rm",
				Description: "remove file of google cloud storage [gs://bucketName/path].",
				ArgsUsage:   "[bucketUri]",
				Action: func(ctx *cli.Context) error {
					CheckEnvironmentValues()
					if ctx.Args().Len() < 1 {
						EprintlnExit("ERROR: rm requires an argument")
						return nil
					}

					uri := ctx.Args().Get(0)
					CommandRemoveObject(uri)
					return nil
				},
			},
			{
				Name:        "rsync",
				Description: "synchronize between google cloud storage files and local files.",
				ArgsUsage:   "[bucketUri|localPath] [localPath|bucketUri]",
				Action: func(ctx *cli.Context) error {
					CheckEnvironmentValues()
					if ctx.Args().Len() < 2 {
						EprintlnExit("ERROR: rsync requires 2 arguments")
						return nil
					}

					srcUri := ctx.Args().Get(0)
					dstUri := ctx.Args().Get(1)
					if strings.HasPrefix(srcUri, "gs://") {
						CommandSyncBucketToLocal(srcUri, dstUri)
					} else {
						CommandSyncLocalToBucket(srcUri, dstUri)
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("ERROR: ", err)
	}
}
