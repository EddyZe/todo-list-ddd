package task

import (
	"context"
	"errors"
	"todo/internal/domain/common"
	"todo/internal/domain/task"
)

func (u *UseCase) ChangeStatus(ctx context.Context, taskID string, status string) error {
	ID, err := common.NewID(taskID)
	if err != nil {
		return err
	}

	t, err := u.repo.GetByID(ctx, ID)
	if err != nil {
		return ErrTaskNotFound
	}

	switch status {
	case task.StatusPending.String():
	case task.StatusInProgress.String():
		if err := t.Start(); err != nil {
			return err
		}
	case task.StatusReview.String():
		if err := t.SendForReview(); err != nil {
			return err
		}
	case task.StatusCancelled.String():
		if err := t.Cancel(); err != nil {
			return err
		}
	case task.StatusDone.String():
		if err := t.Complete(); err != nil {
			return err
		}
	default:
		return errors.New("invalid status")

	}

	if err := u.repo.Save(ctx, t); err != nil {
		return err
	}

	return nil
}
