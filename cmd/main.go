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
		Path: path,
	}); err != nil {
		log.WithError(err).Errorf("unable to apply git log")
		return err
	}
	return nil
}
func main() {
	app := &cli.App{
		Name:  "starter",
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
