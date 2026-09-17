package tasks_dto

type GetTasksInput struct {
	UserID *int
	Limit  *int
	Offset *int
}

type GetTasksOutput []Task
