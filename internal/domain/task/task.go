package task

import (
	"errors"
	"time"
	"todo/internal/domain/comment"
	"todo/internal/domain/common"
)

var (
	ErrAssigneeEmpty    = errors.New("assignee is empty")
	ErrChangeStatusTask = errors.New("failed change status task")
)

type Task struct {
	id          common.ID
	title       Title
	description Description
	creatorID   common.ID
	assigneeID  common.ID
	status      Status
	createdAt   time.Time
	updatedAt   time.Time
	comments    []*comment.Comment
}

func NewTask(title Title, description Description, creatorID common.ID) *Task {
	return &Task{
		id:          common.GenerateID(),
		title:       title,
		description: description,
		creatorID:   creatorID,
		createdAt:   time.Now(),
		updatedAt:   time.Now(),
		status:      StatusPending,
		comments:    make([]*comment.Comment, 0),
	}
}

func (t *Task) ID() common.ID {
	return t.id
}

func (t *Task) Title() Title {
	return t.title
}

func (t *Task) Description() Description {
	return t.description
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Task) Comments() []*comment.Comment {
	return t.comments
}

func (t *Task) SetAssignee(assigneeID common.ID) {
	t.assigneeID = assigneeID
	t.updatedAt = time.Now()
}

func (t *Task) AssigneeID() common.ID {
	return t.assigneeID
}

func (t *Task) CreatorID() common.ID {
	return t.creatorID
}

func (t *Task) Status() Status {
	return t.status
}

func (t *Task) Start() error {
	if t.assigneeID.String() == "" {
		return ErrAssigneeEmpty
	}

	if t.status != StatusPending {
		return errors.Join(ErrChangeStatusTask, errors.New("status is not pending"))
	}

	t.status = StatusInProgress
	t.updatedAt = time.Now()
	return nil
}

func (t *Task) SendForReview() error {
	if t.status != StatusInProgress {
		return errors.Join(ErrChangeStatusTask, errors.New("status is not in-progress"))
	}

	t.status = StatusReview
	t.updatedAt = time.Now()
	return nil
}

func (t *Task) Complete() error {
	if t.status == StatusInProgress || t.status == StatusReview {
		t.status = StatusDone
		t.updatedAt = time.Now()
		return nil
	}

	return errors.Join(ErrChangeStatusTask, errors.New("status is not done"))
}

func (t *Task) Cancel() error {
	if t.status == StatusDone {
		return errors.Join(ErrChangeStatusTask, errors.New("status is done"))
	}

	if t.status == StatusCancelled {
		return errors.Join(ErrChangeStatusTask, errors.New("status is cancelled"))
	}

	t.status = StatusCancelled
	t.updatedAt = time.Now()

	return nil
}

func (t *Task) AddComment(c *comment.Comment) {
	t.comments = append(t.comments, c)
	t.updatedAt = time.Now()
}

func (t *Task) RemoveComment(id common.ID) {
	for i, c := range t.comments {
		if c.ID() == id {
			t.comments = append(t.comments[:i], t.comments[i+1:]...)
			t.updatedAt = time.Now()
			break
		}
	}
}
