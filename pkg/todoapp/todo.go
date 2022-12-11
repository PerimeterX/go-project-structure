package todoapp

import "time"

type Task struct {
	ID    int       `json:"id"`
	Date  time.Time `json:"date"`
	Value string    `json:"value"`
}
