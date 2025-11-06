package task

import (
	"todo/internal/adaptar/http/handlers/gin/user"
	"todo/internal/application/task"
)

// MapTaskToHandler конвертирует task.OutputTask в taskhandler.OutputTask
func MapTaskToHandler(t *task.OutputTask) *OutputTask {
	if t == nil {
		return nil
	}

	comments := make([]*Comment, 0, len(t.Comments))
	for _, c := range t.Comments {
		comments = append(comments, MapCommentToHandler(c))
	}

	o := &OutputTask{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Creator: &user.User{
			ID:    t.Creator.ID,
			Email: t.Creator.Email,
		},
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
		Comments:  comments,
	}

	if t.Assignee != nil {
		o.Assignee = &user.User{
			ID:    t.Assignee.ID,
			Email: t.Assignee.Email,
		}
	}

	return o
}

// MapCommentToHandler конвертирует task.Comment в taskhandler.Comment
func MapCommentToHandler(c *task.Comment) *Comment {
	if c == nil {
		return nil
	}

	return &Comment{
		ID: c.ID,
		Author: &user.User{
			ID:    c.Author.ID,
			Email: c.Author.Email,
		},
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

// MapCreateTaskInput конвертирует DTO запроса в InputCreateTask сервисного слоя
func MapCreateTaskInput(input InputCreateTask, creatorID string) *task.InputCreateTask {
	return &task.InputCreateTask{
		Title:       input.Title,
		Description: input.Description,
		CreatorID:   creatorID,
	}
}

// MapAddCommentInput конвертирует DTO запроса в InputAddComment сервисного слоя
func MapAddCommentInput(input InputAddComment) *task.InputAddComment {
	return &task.InputAddComment{
		Content:  input.Content,
		AuthorID: input.AuthorID,
	}
}

// MapUpdateAssigneeInput конвертирует DTO запроса в InputUpdateAssignee сервисного слоя
func MapUpdateAssigneeInput(input InputUpdateAssignee) *task.InputUpdateAssignee {
	return &task.InputUpdateAssignee{
		TaskID:     input.TaskID,
		AssigneeID: input.AssigneeID,
	}
}

// MapRemoveCommentInput конвертирует DTO запроса в InputRemoveComment сервисного слоя
func MapRemoveCommentInput(input InputRemoveComment) *task.InputRemoveComment {
	return &task.InputRemoveComment{
		TaskID:    input.TaskID,
		CommentID: input.CommentID,
	}
}
