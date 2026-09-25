package statistics_dto

import "github.com/abi-kan/golang-todoapp/internal/core/domain"

type Statistics struct {
	TasksCreatedCount          int      `json:"tasks_created_count"`
	TasksCompletedCount        int      `json:"tasks_completed_count"`
	TasksCompletedRate         *float64 `json:"tasks_completed_rate"`
	TasksAverageCompletionTime *string  `json:"tasks_average_completion_time"`
}

func NewStatisticsDTOFromDomain(statistics domain.Statistics) Statistics {
	var completionTime *string
	if statistics.TasksAverageCompletionTime != nil {
		duration := statistics.TasksAverageCompletionTime.String()
		completionTime = &duration
	}

	return Statistics{
		TasksCreatedCount:          statistics.TasksCreatedCount,
		TasksCompletedCount:        statistics.TasksCompletedCount,
		TasksCompletedRate:         statistics.TasksCompletedRate,
		TasksAverageCompletionTime: completionTime,
	}
}
