package report

import (
	"fmt"
	"os"

	"github.com/saromanov/changelog/pkg/models"
)

type Report interface {
	Do([]models.Message) error
}

// Write provides append changelog data to the file
func Write(filename, data string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("unable to open file %s: %v", filename, err)
	}

	defer f.Close()

	if _, err = f.WriteString(data); err != nil {
		return fmt.Errorf("unable to write to file: %v", err)
	}
	return nil
}
