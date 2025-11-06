package task

import (
	"context"
	"todo/internal/domain/common"
)

func (u *UseCase) DeleteComment(ctx context.Context, comment InputRemoveComment) error {
	taskID, err := common.NewID(comment.TaskID)
	if err != nil {
		return err
	}

	commentID, err := common.NewID(comment.CommentID)
	if err != nil {
		return err
	}

	t, err := u.repo.GetByID(ctx, taskID)
	if err != nil {
		return ErrTaskNotFound
	}

	t.RemoveComment(commentID)
	if err := u.repo.Save(ctx, t); err != nil {
		return err
	}

	return nil
}
