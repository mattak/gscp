package main

import (
	"errors"
	"fmt"
	"github.com/mattak/gscp/pkg/gscp"
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
		Name:        "gscp",
		Usage:       "minimum set of google cloud copy command",
		Description: "Environment value of GOOGLE_APPLICATION_CREDENTIALS is required for every command processing",
		Commands: []*cli.Command{
			{
				Name:        "ls",
				Description: "list up objects of google cloud storage.",
				ArgsUsage:   "[bucketUri:\"gs://bucketName/path\"]",
				Action: func(ctx *cli.Context) error {
					gscp.CheckEnvironmentValues()

					if ctx.Args().Len() < 1 {
						EprintlnExit("ERROR: ls requires a argument")
						return nil
					}

					uri := ctx.Args().First()
					return gscp.CommandListObjects(uri)
				},
			},
			{
				Name:        "lsd",
				Description: "list up detail objects of google cloud storage.",
				ArgsUsage:   "[bucketUri:\"gs://bucketName/path\"]",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "latest1",
						Aliases: []string{"1"},
					},
					&cli.BoolFlag{
						Name:    "name-only",
						Aliases: []string{"n"},
					},
					&cli.BoolFlag{
						Name:    "unix-millis",
						Aliases: []string{"u"},
					},
				},
				Action: func(ctx *cli.Context) error {
					gscp.CheckEnvironmentValues()

					if ctx.Args().Len() < 1 {
						EprintlnExit("ERROR: ls requires a argument")
						return nil
					}

					uri := ctx.Args().First()
					withLatest1Option := ctx.Bool("latest1")
					withNameOnlyOption := ctx.Bool("name-only")
					withUnixMillisOption := ctx.Bool("unix-millis")
					option := gscp.CommandListDetailObjectsOption{
						WithLatest1:    withLatest1Option,
						WithNameOnly:   withNameOnlyOption,
						WithUnixMillis: withUnixMillisOption,
					}
					return gscp.CommandListDetailObjects(uri, option)
				},
			},
			{
				Name:        "cp",
				Description: "copy files between google cloud storage files [gs://bucketName/path] and local files [/path/to/file].",
				ArgsUsage:   "[bucketUri|localPath] [localPath|bucketUri]",
				Action: func(ctx *cli.Context) error {
					gscp.CheckEnvironmentValues()

					if ctx.Args().Len() < 2 {
						EprintlnExit("ERROR: cp requires 2 arguments")
						return nil
					}

					srcUri := ctx.Args().Get(0)
					dstUri := ctx.Args().Get(1)
					if strings.HasPrefix(srcUri, "gs://") {
						return gscp.CommandCopyObjectToLocal(srcUri, dstUri)
					} else {
						return gscp.CommandCopyLocalToBucket(srcUri, dstUri)
					}
				},
			},
			{
				Name:        "rm",
				Description: "remove file of google cloud storage [gs://bucketName/path].",
				ArgsUsage:   "[bucketUri]",
				Action: func(ctx *cli.Context) error {
					gscp.CheckEnvironmentValues()
					if ctx.Args().Len() < 1 {
						EprintlnExit("ERROR: rm requires an argument")
						return nil
					}

					uri := ctx.Args().Get(0)
					return gscp.CommandRemoveObject(uri)
				},
			},
			{
				Name:        "rsync",
				Description: "synchronize between google cloud storage files and local files.",
				ArgsUsage:   "[bucketUri|localPath] [localPath|bucketUri]",
				Action: func(ctx *cli.Context) error {
					gscp.CheckEnvironmentValues()
					if ctx.Args().Len() < 2 {
						EprintlnExit("ERROR: rsync requires 2 arguments")
						return nil
					}

					srcUri := ctx.Args().Get(0)
					dstUri := ctx.Args().Get(1)
					if strings.HasPrefix(srcUri, "gs://") {
						return gscp.CommandSyncBucketToLocal(srcUri, dstUri)
					} else {
						return gscp.CommandSyncLocalToBucket(srcUri, dstUri)
					}
				},
			},
			{
				Name:        "watch",
				Description: "watch local dir to copy google cloud storage files.",
				ArgsUsage:   "[localPath] [bucketUriDir]",
				Action: func(ctx *cli.Context) error {
					gscp.CheckEnvironmentValues()
					if ctx.Args().Len() < 2 {
						EprintlnExit("ERROR: rsync requires 2 arguments")
						return nil
					}

					localDir := ctx.Args().Get(0)
					bucketUri := ctx.Args().Get(1)
					if strings.HasPrefix(bucketUri, "gs://") {
						return gscp.CommandWatchSync(localDir, bucketUri)
					} else {
						return errors.New("Google Cloud Storage uri should be start with gs:// prefix.")
					}
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal("ERROR: ", err)
	}
}
