package core

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// ReleaseRequest defines request for making release notes
type ReleaseRequest struct {
	Path string
}

// MakeReleaseNotes provides creating of release notes based on git commits
func MakeReleaseNotes(r ReleaseRequest) error {
	messages, err := log(r.Path, func(m string) bool {
		return len(m) > 0
	})
	if err != nil {
		return fmt.Errorf("unable to get log messages")
	}
	fmt.Println(messages)
	return nil
}

// log returns git log
func log(path string, filter func(string) bool) ([]string, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open repo: %v", err)
	}
	cIter, err := r.Log(&git.LogOptions{})
	if err != nil {
		return nil, fmt.Errorf("unable to execute git log: %v", err)
	}

	commits := []string{}
	err = cIter.ForEach(func(c *object.Commit) error {
		if filter(c.Message) {
			commits = append(commits, c.Message)
		}

		return nil
	})

	return commits, nil
}
