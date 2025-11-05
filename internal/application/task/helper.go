package task

import (
	"todo/internal/application/user"
	"todo/internal/domain/common"
	"todo/internal/domain/task"
)

func mapTaskToOutputTask(t *task.Task, getUser func(id common.ID) (*user.OutputUser, error)) (*OutputTask, error) {
	if t == nil {
		return nil, nil
	}

	comments := make([]*Comment, 0, len(t.Comments()))
	for _, c := range t.Comments() {
		author, err := getUser(c.AuthorID())
		if err != nil {
			return nil, err
		}

		comments = append(comments, &Comment{
			ID:        c.ID().String(),
			Author:    author,
			Content:   c.Content().String(),
			CreatedAt: c.CreatedAt(),
			UpdatedAt: c.UpdatedAt(),
		})
	}

	creator, err := getUser(t.CreatorID())
	if err != nil {
		return nil, err
	}

	var assignee *user.OutputUser
	if t.AssigneeID().String() != "" {
		assignee, err = getUser(t.AssigneeID())
		if err != nil {
			return nil, err
		}
	}

	return &OutputTask{
		ID:          t.ID().String(),
		Title:       t.Title().String(),
		Description: t.Description().String(),
		Creator:     creator,
		Assignee:    assignee,
		Status:      t.Status().String(),
		CreatedAt:   t.CreatedAt(),
		UpdatedAt:   t.UpdatedAt(),
		Comments:    comments,
	}, nil
}
