package main

import (
	"os"

	"github.com/saromanov/changelog/pkg/core"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	defaultPath = "."
)

func build(c *cli.Context) error {
	path := c.String("path")
	if path == "" {
		path = defaultPath
	}

	if err := core.MakeReleaseNotes(core.ReleaseRequest{
		Path:     path,
		Filename: c.String("filename"),
		Type:     c.String("type"),
		Title:    c.String("title"),
		Pattern:  c.String("pattern"),
	}); err != nil {
		log.WithError(err).Errorf("unable to apply git log")
		return err
	}
	return nil
}
func main() {
	app := &cli.App{
		Name:  "changelog",
		Usage: "create puppet for the project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "path",
				Value: "",
				Usage: "path to git directory",
			},
			&cli.StringFlag{
				Name:  "since",
				Value: "",
				Usage: "return commits from the date",
			},
			&cli.StringFlag{
				Name:  "until",
				Value: "",
				Usage: "return commits until the date",
			},
			&cli.StringFlag{
				Name:  "version",
				Value: "",
				Usage: "generated new version",
			},
			&cli.StringFlag{
				Name:  "type",
				Value: "",
				Usage: "type of the output",
			},
			&cli.StringFlag{
				Name:  "filename",
				Value: "",
				Usage: "filename of the output",
			},
			&cli.StringFlag{
				Name:  "title",
				Value: "",
				Usage: "title of the release",
			},
			&cli.StringFlag{
				Name:  "pattern",
				Value: "",
				Usage: "pattern for get commit for changelog",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "build",
				Usage:  "building of the new project",
				Action: build,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		return
	}
}
