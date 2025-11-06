package task

import (
	"time"
	"todo/internal/adaptar/http/handlers/gin/user"
)

type ListResponse struct {
	Tasks []*OutputTask `json:"tasks"`
}

type InputCreateTask struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Comment struct {
	ID        string     `json:"id"`
	Author    *user.User `json:"author"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type InputAddComment struct {
	Content  string `json:"content"`
	AuthorID string `json:"author_id,omitempty"`
}

type InputRemoveComment struct {
	TaskID    string `json:"task_id"`
	CommentID string `json:"comment_id"`
}

type InputUpdateAssignee struct {
	TaskID     string `json:"task_id"`
	AssigneeID string `json:"assignee_id"`
}

type User struct {
}

type OutputTask struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Creator     *user.User `json:"creator"`
	Assignee    *user.User `json:"assignee"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Comments    []*Comment `json:"comments"`
}
