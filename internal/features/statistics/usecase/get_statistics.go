package statistics_usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/abi-kan/golang-todoapp/internal/core/domain"
	core_errors "github.com/abi-kan/golang-todoapp/internal/core/errors"
	statistics_dto "github.com/abi-kan/golang-todoapp/internal/features/statistics/dto"
)

func (s *Statistics) GetStatistics(
	ctx context.Context,
	input statistics_dto.GetStatisticsInput,
) (statistics_dto.GetStatisticsOutput, error) {
	if input.From != nil && input.To != nil {
		if input.To.Before(*input.From) || input.To.Equal(*input.From) {
			return statistics_dto.GetStatisticsOutput{}, fmt.Errorf(
				"'to' must be after 'from': %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	tasks, err := s.postgres.GetTasks(
		ctx, input.UserID, input.From, input.To,
	)
	if err != nil {
		return statistics_dto.GetStatisticsOutput{}, fmt.Errorf(
			"get tasks from repository: %w",
			err,
		)
	}

	output := statistics_dto.GetStatisticsOutput(
		calculateStatistics(tasks),
	)

	return output, nil
}

func calculateStatistics(tasks []domain.Task) statistics_dto.Statistics {
	if len(tasks) == 0 {
		return statistics_dto.Statistics{}
	}

	createdCount := len(tasks)
	completedCount := 0
	var totalCompletionDuration time.Duration
	for _, task := range tasks {
		if task.Completed {
			completedCount++
		}

		if completionDuration := task.CompletionDuration(); completionDuration != nil {
			totalCompletionDuration += *completionDuration
		}
	}

	completedRate := float64(completedCount) / float64(createdCount) * 100
	var averageCompletionTime *time.Duration
	if completedCount > 0 && totalCompletionDuration != 0 {
		avg := totalCompletionDuration / time.Duration(completedCount)
		averageCompletionTime = &avg
	}

	return statistics_dto.NewStatisticsDTOFromDomain(
		domain.NewStatistics(
			createdCount,
			completedCount,
			&completedRate,
			averageCompletionTime,
		),
	)
}
