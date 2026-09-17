package tasks_usecase

import (
	"context"
	"fmt"
)

func (t *Tasks) DeleteTask(
	ctx context.Context,
	id int,
) error {
	if err := t.postgres.DeleteTask(ctx, id); err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	return nil
}
