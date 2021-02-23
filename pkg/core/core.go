package core

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/saromanov/changelog/pkg/models"
	"github.com/saromanov/changelog/pkg/report"
	"github.com/saromanov/changelog/pkg/report/txt"
)

const defaultReleaseTitle = "New Release!"

// ReleaseRequest defines request for making release notes
type ReleaseRequest struct {
	Path     string
	Type     string
	Filename string
	Title    string
	Since    time.Time
	Until    time.Time
}

// MakeReleaseNotes provides creating of release notes based on git commits
func MakeReleaseNotes(r ReleaseRequest) error {
	if r.Filename == "" {
		return fmt.Errorf("filename is not defined")
	}
	r.Title = makeReleaseTitle(r.Title)
	messages, err := log(r, func(m string) bool {
		return len(m) > 0
	})
	if err != nil {
		return fmt.Errorf("unable to get log messages")
	}
	return makeOutput(r, messages)
}

// log returns git log
func log(req ReleaseRequest, filter func(string) bool) ([]models.Message, error) {
	r, err := git.PlainOpen(req.Path)
	if err != nil {
		return nil, fmt.Errorf("unable to open repo: %v", err)
	}
	opt := &git.LogOptions{}
	if !req.Since.IsZero() {
		opt.Since = &req.Since
	}
	if !req.Until.IsZero() {
		opt.Until = &req.Until
	}
	cIter, err := r.Log(opt)
	if err != nil {
		return nil, fmt.Errorf("unable to execute git log: %v", err)
	}

	tags, err := r.TagObjects()
	if err != nil {
		return nil, fmt.Errorf("unable to get tags: %v", err)
	}
	tags.ForEach(func(c *object.Tag) error {
		fmt.Println(c.Tagger.When.Clock())
		return nil
	})
	commits := []models.Message{}
	err = cIter.ForEach(func(c *object.Commit) error {
		if filter(c.Message) {
			commits = append(commits, models.Message{Message: strings.Trim(c.Message, "\n"), Author: c.Author.Name, Date: c.Author.When})
		}
		return nil
	})

	return commits, nil
}

// returns prepared title for release
func makeReleaseTitle(title string) string {
	if title == "" {
		title = defaultReleaseTitle
	}
	return fmt.Sprintf("%s(%s)\n", title, time.Now().Format(time.RFC3339))
}

func makeOutput(r ReleaseRequest, m []models.Message) error {
	d := func(rr ReleaseRequest) report.Report {
		return txt.New(rr.Filename, rr.Title)
	}
	return d(r).Do(m)
}
