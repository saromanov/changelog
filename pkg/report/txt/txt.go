package txt

import (
	"fmt"
	"os"
	"time"

	"github.com/saromanov/changelog/pkg/models"
	"github.com/saromanov/changelog/pkg/report"
)

// txt defines txt report
type txt struct {
	filename string
}

// New provides initialization of txt report
func New(filename string) report.Report {
	return &txt{
		filename: filename,
	}
}

// Do provides generation of the report
func (t *txt) Do(data []models.Message) error {
	if len(data) == 0 {
		return nil
	}

	var result string
	for _, d := range data {
		result += fmt.Sprintf("%s %s (%s)\n", d.Date.Format(time.RFC3339), d.Message, d.Author)
	}
	return write(t.filename, result)
}

// write provides append changelog data to the file
func write(filename, data string) error {
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
