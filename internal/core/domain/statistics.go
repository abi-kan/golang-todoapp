package domain

import "time"

type Statistics struct {
	TasksCreatedCount          int
	TasksCompletedCount        int
	TasksCompletedRate         *float64
	TasksAverageCompletionTime *time.Duration
}

func NewStatistics(
	tasksCreatedCount int,
	tasksCompletedCount int,
	tasksCompletedRate *float64,
	tasksAverageCompletionTime *time.Duration,
) Statistics {
	return Statistics{
		TasksCreatedCount:          tasksCreatedCount,
		TasksCompletedCount:        tasksCompletedCount,
		TasksCompletedRate:         tasksCompletedRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}
