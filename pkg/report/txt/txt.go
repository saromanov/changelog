package txt

import (
	"fmt"
	"time"

	"github.com/saromanov/changelog/pkg/models"
	"github.com/saromanov/changelog/pkg/report"
)

// txt defines report for txt
type txt struct {
	filename string
	title    string
}

// New provides initialization of txt report
func New(filename, title string) report.Report {
	return &txt{
		filename: filename,
		title:    title,
	}
}

// Do provides generation of the report
func (t *txt) Do(data []models.Message) error {
	if len(data) == 0 {
		return nil
	}

	result := t.title
	for _, d := range data {
		result += fmt.Sprintf("%s %s (%s)\n", d.Date.Format(time.RFC3339), d.Message, d.Author)
	}
	result += "\n"
	return report.Write(t.filename, result)
}
