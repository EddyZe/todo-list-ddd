package comment_test

import (
	"testing"
	"time"
	"todo/internal/domain/comment"
	"todo/internal/domain/common"
)

func TestCommentAggregate(t *testing.T) {
	authorID := common.GenerateID()
	taskID := common.GenerateID()

	t.Run("create comment", func(t *testing.T) {
		content, _ := comment.NewContent("Hello World")
		c := comment.NewComment(authorID, taskID, content)

		if c.AuthorID() != authorID {
			t.Errorf("expected authorID %v, got %v", authorID, c.AuthorID())
		}

		if c.TaskID() != taskID {
			t.Errorf("expected taskID %v, got %v", taskID, c.TaskID())
		}

		if c.Content() != content {
			t.Errorf("expected content %v, got %v", content, c.Content())
		}

		if time.Since(c.CreatedAt()) > time.Second {
			t.Errorf("createdAt seems wrong")
		}

		if time.Since(c.UpdatedAt()) > time.Second {
			t.Errorf("updatedAt seems wrong")
		}
	})

	t.Run("change content", func(t *testing.T) {
		content, _ := comment.NewContent("Original")
		c := comment.NewComment(authorID, taskID, content)

		newContent, _ := comment.NewContent("Updated Content")
		oldUpdatedAt := c.UpdatedAt()
		time.Sleep(time.Millisecond * 10)
		c.ChangeContent(newContent)

		if c.Content() != newContent {
			t.Errorf("expected content %v, got %v", newContent, c.Content())
		}

		if !c.UpdatedAt().After(oldUpdatedAt) {
			t.Errorf("expected updatedAt to be refreshed")
		}
	})
}
