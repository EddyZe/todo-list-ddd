package task

import (
	"context"
	"errors"
	"todo/internal/domain/common"
	"todo/internal/domain/task"
)


func (u *UseCase) UpdateAssignee(ctx context.Context, input InputUpdateAssignee) error {
	taskID, err := common.NewID(input.TaskID)
	if err != nil {
		return errors.Join(ErrUpdateAssignee, err)
	}

	assigneeID, err := common.NewID(input.AssigneeID)
	if err != nil {
		return errors.Join(ErrUpdateAssignee, err)
	}

	t, err := u.repo.GetByID(ctx, taskID)
	if err != nil {
		if errors.Is(task.ErrTaskNotFound, err) {
			return ErrTaskNotFound
		}
		return errors.Join(ErrUpdateAssignee, err)
	}

	if _, err := u.userReader.GetByID(ctx, input.AssigneeID); err != nil {
		return errors.Join(ErrUpdateAssignee, err)
	}

	t.SetAssignee(assigneeID)

	if err := u.repo.Save(ctx, t); err != nil {
		return errors.Join(ErrUpdateAssignee, err)
	}

	return nil
}
