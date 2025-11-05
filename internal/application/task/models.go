package task

import (
	"time"
	"todo/internal/application/user"
)

type InputCreateTask struct {
	Title       string
	Description string
	CreatorID   string
}

type Comment struct {
	ID        string
	Author    *user.OutputUser
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InputAddComment struct {
	Content  string
	AuthorID string
}

type InputRemoveComment struct {
	TaskID    string
	CommentID string
}

type InputUpdateAssignee struct {
	TaskID     string
	AssigneeID string
}

type OutputTask struct {
	ID          string
	Title       string
	Description string
	Creator     *user.OutputUser
	Assignee    *user.OutputUser
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Comments    []*Comment
}
