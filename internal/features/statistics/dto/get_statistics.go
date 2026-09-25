package statistics_dto

import "time"

type GetStatisticsInput struct {
	UserID *int
	From   *time.Time
	To     *time.Time
}

type GetStatisticsOutput Statistics
