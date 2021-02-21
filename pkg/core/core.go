package core

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// ReleaseRequest defines request for making release notes
type ReleaseRequest struct {
	Path  string
	Since time.Time
	Until time.Time
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
func log(req ReleaseRequest, filter func(string) bool) ([]Report, error) {
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

	tags, err := r.Tags()
	if err != nil {
		return nil, fmt.Errorf("unable to get tags: %v", err)
	}

	tags.ForEach(func(c *plumbing.Reference) error {
		fmt.Println(c.Name())
		return nil
	})
	commits := []Report{}
	err = cIter.ForEach(func(c *object.Commit) error {
		if filter(c.Message) {
			commits = append(commits, Report{Message: c.Message, Author: c.Author.Name})
		}

		return nil
	})

	return commits, nil
}
