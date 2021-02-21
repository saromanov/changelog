package core

import "time"

// Report defines struct for generated report
type Report struct {
	Message string    `json:"message"`
	Author  string    `json:"author"`
	Date    time.Time `json:"time"`
}
