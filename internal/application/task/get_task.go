package task

import (
	"context"
	"errors"
	"todo/internal/application/user"
	"todo/internal/domain/common"
	"todo/internal/domain/task"
)

func (u *UseCase) GetByID(ctx context.Context, taskID string) (*OutputTask, error) {
	id, err := common.NewID(taskID)
	if err != nil {
		return nil, errors.New("invalid task id")
	}

	t, err := u.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, task.ErrTaskNotFound) {
			return nil, ErrTaskNotFound
		}
		return nil, errors.New("failed getting task")
	}

	output, err := mapTaskToOutputTask(t, func(id common.ID) (*user.OutputUser, error) {
		return u.userReader.GetByID(ctx, id.String())
	})
	if err != nil {
		return nil, errors.New("failed getting task")
	}

	return output, nil
}

func (u *UseCase) GetByAuthorID(ctx context.Context, authorID string) []*OutputTask {
	return u.findTasks(ctx, authorID, u.repo.GetByAuthorID)
}

func (u *UseCase) GetByAssigneeID(ctx context.Context, assigneeID string) []*OutputTask {
	return u.findTasks(ctx, assigneeID, u.repo.GetByAssigneeID)
}

func (u *UseCase) findTasks(ctx context.Context, id string, f func(ctx context.Context, id common.ID) ([]*task.Task, error)) []*OutputTask {
	comID, err := common.NewID(id)
	if err != nil {
		return make([]*OutputTask, 0)
	}

	ts, err := f(ctx, comID)
	if err != nil {
		return make([]*OutputTask, 0)
	}

	res := make([]*OutputTask, 0, len(ts))
	for _, t := range ts {
		o, err := mapTaskToOutputTask(t, func(id common.ID) (*user.OutputUser, error) {
			return u.userReader.GetByID(ctx, id.String())
		})
		if err != nil {
			continue
		}
		res = append(res, o)
	}
	return res
}
