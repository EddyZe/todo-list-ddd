package task

import (
	"context"
	"todo/internal/application/user"
	"todo/internal/domain/common"
)

func (u *UseCase) GetAll(ctx context.Context) []*OutputTask {
	tasks, err := u.repo.GetAll(ctx)
	if err != nil {
		return make([]*OutputTask, 0)
	}

	result := make([]*OutputTask, 0, len(tasks))

	for _, t := range tasks {
		o, err := mapTaskToOutputTask(t, func(id common.ID) (*user.OutputUser, error) {
			return u.userReader.GetByID(ctx, id.String())
		})
		if err != nil {
			continue
		}

		result = append(result, o)
	}

	return result
}
