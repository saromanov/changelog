package core

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/saromanov/changelog/pkg/models"
)

// ReleaseRequest defines request for making release notes
type ReleaseRequest struct {
	Path     string
	Type     string
	Filename string
	Since    time.Time
	Until    time.Time
}

// MakeReleaseNotes provides creating of release notes based on git commits
func MakeReleaseNotes(r ReleaseRequest) error {
	messages, err := log(r, func(m string) bool {
		return len(m) > 0
	})
	if err != nil {
		return fmt.Errorf("unable to get log messages")
	}
	fmt.Println(messages)
	return nil
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
			commits = append(commits, models.Message{Message: c.Message, Author: c.Author.Name, Date: c.Author.When})
		}
		return nil
	})

	return commits, nil
}
