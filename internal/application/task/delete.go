package task

import (
	"context"
	"errors"
	"todo/internal/domain/common"
	"todo/internal/domain/task"
)

func (u *UseCase) DeleteByID(ctx context.Context, taskID string) error {
	id, err := common.NewID(taskID)
	if err != nil {
		return errors.Join(ErrDeletingTask, err)
	}

	if err := u.repo.DeleteByID(ctx, id); err != nil {
		if errors.Is(err, task.ErrTaskNotFound) {
			return errors.Join(ErrTaskNotFound, err)
		}

		return errors.Join(ErrDeletingTask, err)
	}

	return nil
}
