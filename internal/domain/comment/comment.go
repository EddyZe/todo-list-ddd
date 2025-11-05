package comment

import (
	"time"
	"todo/internal/domain/common"
)

type Comment struct {
	id        common.ID
	authorID  common.ID
	taskID    common.ID
	content   Content
	createdAt time.Time
	updatedAt time.Time
}

func NewComment(authorID common.ID, taskID common.ID, content Content) *Comment {
	return &Comment{
		id:        common.GenerateID(),
		authorID:  authorID,
		taskID:    taskID,
		content:   content,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (c *Comment) ID() common.ID {
	return c.id
}

func (c *Comment) AuthorID() common.ID {
	return c.authorID
}

func (c *Comment) TaskID() common.ID {
	return c.taskID
}

func (c *Comment) Content() Content {
	return c.content
}

func (c *Comment) ChangeContent(content Content) {
	c.content = content
	c.updatedAt = time.Now()
}

func (c *Comment) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Comment) UpdatedAt() time.Time {
	return c.updatedAt
}
