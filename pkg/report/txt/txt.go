package txt

import (
	"fmt"
	"io/ioutil"

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
		fmt.Println(d)
	}
	return ioutil.WriteFile(t.filename, []byte(result), 0644)
}
