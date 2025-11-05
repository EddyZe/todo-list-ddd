package task

import "errors"

var (
	ErrUpdateAssignee = errors.New("failed update assignee")
	ErrTaskNotFound   = errors.New("task not found")
	ErrCreatingTask   = errors.New("failed creating task")
	ErrDeletingTask   = errors.New("failed deleting task")
)
