package report

import "github.com/saromanov/changelog/pkg/models"
type Report interface {
	Do([]models.Message) error
}