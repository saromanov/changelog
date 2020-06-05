package core

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Log returns git log
func Log(path string) error {
	r, err := git.PlainOpen(path)
	if err != nil {
		return fmt.Errorf("unable to open repo: %v", err)
	}
	cIter, err := r.Log(&git.LogOptions{})
	if err != nil {
		return fmt.Errorf("unable to execute git log: %v", err)
	}

	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)

		return nil
	})

	return nil
}
