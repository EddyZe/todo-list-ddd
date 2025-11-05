package task

import (
	"testing"
	"time"
	"todo/internal/domain/comment"
	"todo/internal/domain/common"
	"todo/internal/domain/task"
)

func TestTaskAggregate(t *testing.T) {
	title, err := task.NewTitle("Test Task")
	if err != nil {
		t.Fatalf("error creating title: %v", err)
	}
	desc, err := task.NewDescription("Task description")
	if err != nil {
		t.Fatalf("error creating description: %v", err)
	}
	creatorID := common.GenerateID()
	assigneeID := common.GenerateID()

	t.Run("create task", func(t *testing.T) {
		taskObj := task.NewTask(title, desc, creatorID)

		if taskObj.CreatorID() != creatorID {
			t.Errorf("expected creatorID %v, got %v", creatorID, taskObj.CreatorID())
		}

		if taskObj.Title() != title {
			t.Errorf("expected title %v, got %v", title, taskObj.Title())
		}

		if taskObj.Description() != desc {
			t.Errorf("expected description %v, got %v", desc, taskObj.Description())
		}

		if taskObj.Status() != task.StatusPending {
			t.Errorf("expected status PENDING, got %v", taskObj.Status())
		}
	})

	t.Run("assign and start task", func(t *testing.T) {
		taskObj := task.NewTask(title, desc, creatorID)
		taskObj.SetAssignee(assigneeID)

		if err := taskObj.Start(); err != nil {
			t.Fatalf("unexpected error starting task: %v", err)
		}

		if taskObj.Status() != task.StatusInProgress {
			t.Errorf("expected status IN_PROGRESS, got %v", taskObj.Status())
		}
	})

	t.Run("send for review and complete task", func(t *testing.T) {
		taskObj := task.NewTask(title, desc, creatorID)
		taskObj.SetAssignee(assigneeID)
		err = taskObj.Start()
		if err != nil {
			t.Fatalf("unexpected error starting task: %v", err)
		}

		if err := taskObj.SendForReview(); err != nil {
			t.Fatalf("unexpected error sending for review: %v", err)
		}
		if taskObj.Status() != task.StatusReview {
			t.Errorf("expected status REVIEW, got %v", taskObj.Status())
		}

		if err := taskObj.Complete(); err != nil {
			t.Fatalf("unexpected error completing task: %v", err)
		}
		if taskObj.Status() != task.StatusDone {
			t.Errorf("expected status DONE, got %v", taskObj.Status())
		}
	})

	t.Run("cancel task", func(t *testing.T) {
		taskObj := task.NewTask(title, desc, creatorID)
		taskObj.SetAssignee(assigneeID)
		err = taskObj.Start()
		if err != nil {
			t.Fatalf("unexpected error starting task: %v", err)
		}

		if err := taskObj.Cancel(); err != nil {
			t.Fatalf("unexpected error cancelling task: %v", err)
		}
		if taskObj.Status() != task.StatusCancelled {
			t.Errorf("expected status CANCELLED, got %v", taskObj.Status())
		}
	})

	t.Run("add and remove comment", func(t *testing.T) {
		taskObj := task.NewTask(title, desc, creatorID)
		c := comment.NewComment(creatorID, taskObj.ID(), comment.Content("My comment"))
		taskObj.AddComment(c)

		if len(taskObj.Comments()) != 1 {
			t.Errorf("expected 1 comment, got %d", len(taskObj.Comments()))
		}

		taskObj.RemoveComment(c.ID())
		if len(taskObj.Comments()) != 0 {
			t.Errorf("expected 0 comments after removal, got %d", len(taskObj.Comments()))
		}
	})

	t.Run("updatedAt changes", func(t *testing.T) {
		taskObj := task.NewTask(title, desc, creatorID)
		oldUpdated := taskObj.UpdatedAt()
		time.Sleep(time.Millisecond * 10)

		taskObj.SetAssignee(assigneeID)
		if !taskObj.UpdatedAt().After(oldUpdated) {
			t.Errorf("expected updatedAt to be refreshed")
		}
	})
}
