package tasks

import "time"

type Task struct {
	ID int `json:"id"`
	Description string `json:"description"`
	Status string `"json:"status"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

func (a *Task) Update(description string) *Task {
	a.Description = description
	return a
}