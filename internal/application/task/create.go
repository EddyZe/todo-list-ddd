package task

import (
	"context"
	"errors"
	"todo/internal/application/user"
	"todo/internal/domain/common"
	"todo/internal/domain/task"
)

func (u *UseCase) CreateTask(ctx context.Context, input InputCreateTask) (output *OutputTask, err error) {
	title, err := task.NewTitle(input.Title)
	if err != nil {
		return nil, errors.Join(ErrCreatingTask, err)
	}

	description, err := task.NewDescription(input.Description)
	if err != nil {
		return nil, errors.Join(ErrCreatingTask, err)
	}

	creatorID, err := common.NewID(input.CreatorID)
	if err != nil {
		return nil, errors.Join(ErrCreatingTask, err)
	}

	if _, err := u.userReader.GetByID(ctx, creatorID.String()); err != nil {
		return nil, errors.Join(ErrCreatingTask, err)
	}

	newTask := task.NewTask(title, description, creatorID)

	if err := u.repo.Save(ctx, newTask); err != nil {
		return nil, errors.Join(ErrCreatingTask, err)
	}

	output, err = mapTaskToOutputTask(newTask, func(id common.ID) (*user.OutputUser, error) {
		return u.userReader.GetByID(ctx, id.String())
	})
	if err != nil {
		return nil, errors.Join(ErrCreatingTask, err)
	}

	return output, nil
}
