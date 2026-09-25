package tasks_dto

import (
	"time"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
)

type Task struct {
	ID          int        `json:"id"`
	Version     int        `json:"version"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`

	AuthorUserID int `json:"author_user_id"`
}

func NewTaskDTOFromDomain(task domain.Task) Task {
	return Task{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func NewTasksDTOFromDomain(tasks []domain.Task) []Task {
	dtoTasks := make([]Task, len(tasks))
	for index, task := range tasks {
		dtoTasks[index] = NewTaskDTOFromDomain(task)
	}

	return dtoTasks
}
