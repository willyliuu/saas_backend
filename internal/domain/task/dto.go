package task

import "time"

type CreateTaskDTO struct {
	Title       string
	Description string
	Status      string
	DueDate     time.Time
	AssigneeID  *uint
}

type UpdateTaskDTO struct {
	Title       string
	Description string
	Status      string
	DueDate     time.Time
	AssigneeID  *uint
}
