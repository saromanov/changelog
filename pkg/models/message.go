package models

import "time"

// Message defines struct for generated report
type Message struct {
	Message string    `json:"message"`
	Author  string    `json:"author"`
	Date    time.Time `json:"time"`
}
