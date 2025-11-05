package task

import (
	"context"
	"todo/internal/application/user"
	"todo/internal/domain/comment"
	"todo/internal/domain/common"
)

func (u *UseCase) AddComment(ctx context.Context, taskID string, c InputAddComment) error {
	id, err := common.NewID(taskID)
	if err != nil {
		return err
	}

	t, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if _, err := u.userReader.GetByID(ctx, c.AuthorID); err != nil {
		return user.ErrUserNotFound
	}

	authorID, err := common.NewID(c.AuthorID)
	if err != nil {
		return err
	}

	content, err := comment.NewContent(c.Content)
	if err != nil {
		return err
	}

	newComment := comment.NewComment(authorID, t.ID(), content)

	t.AddComment(newComment)
	if err := u.repo.Save(ctx, t); err != nil {
		return err
	}

	return nil
}
