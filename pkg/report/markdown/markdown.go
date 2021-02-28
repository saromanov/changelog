package markdown

import (
	"fmt"
	"time"

	"github.com/saromanov/changelog/pkg/models"
	"github.com/saromanov/changelog/pkg/report"
)

// markdown defines report for markdown
type markdown struct {
	filename string
	title    string
}

// New provides initialization of markdown report
func New(filename, title string) report.Report {
	return &markdown{
		filename: filename,
		title:    title,
	}
}

// Do provides generation of the report
func (t *markdown) Do(data []models.Message) error {
	if len(data) == 0 {
		return nil
	}

	result := "# " + t.title
	for _, d := range data {
		result += fmt.Sprintf("* %s %s (%s)\n", d.Date.Format(time.RFC3339), d.Message, d.Author)
	}
	result += "\n"
	return report.Write(t.filename, result)
}
