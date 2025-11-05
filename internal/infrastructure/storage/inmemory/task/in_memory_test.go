package task

import (
	"context"
	"testing"
	"todo/internal/domain/common"
	"todo/internal/domain/task"
)

func TestInMemory_SaveAndGetByID(t *testing.T) {
	repo := NewInMemory()
	ctx := context.Background()

	title, _ := task.NewTitle("Test Task")
	desc, _ := task.NewDescription("Testing save and get")
	tk := task.NewTask(title, desc, common.GenerateID())

	err := repo.Save(ctx, tk)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got, err := repo.GetByID(ctx, tk.ID())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got.ID() != tk.ID() {
		t.Errorf("expected String %v, got %v", tk.ID(), got.ID())
	}
	if got.Title() != tk.Title() {
		t.Errorf("expected title %v, got %v", tk.Title(), got.Title())
	}
}

func TestInMemory_DeleteByID(t *testing.T) {
	repo := NewInMemory()
	ctx := context.Background()

	title, _ := task.NewTitle("To Delete")
	desc, _ := task.NewDescription("For delete test")
	tk := task.NewTask(title, desc, common.GenerateID())

	_ = repo.Save(ctx, tk)

	err := repo.DeleteByID(ctx, tk.ID())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = repo.GetByID(ctx, tk.ID())
	if err == nil {
		t.Fatal("expected error after delete, got nil")
	}
}

func TestInMemory_GetByAuthorID(t *testing.T) {
	repo := NewInMemory()
	ctx := context.Background()

	authorID := common.GenerateID()
	otherAuthor := common.GenerateID()

	t1 := task.NewTask(task.Title("Task 1"), task.Description("A"), authorID)
	t2 := task.NewTask(task.Title("Task 2"), task.Description("B"), authorID)
	t3 := task.NewTask(task.Title("Task 3"), task.Description("C"), otherAuthor)

	_ = repo.Save(ctx, t1)
	_ = repo.Save(ctx, t2)
	_ = repo.Save(ctx, t3)

	list, err := repo.GetByAuthorID(ctx, authorID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(list) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(list))
	}
}

func TestInMemory_GetByAssigneeID(t *testing.T) {
	repo := NewInMemory()
	ctx := context.Background()

	assigneeID := common.GenerateID()
	otherAssignee := common.GenerateID()

	t1 := task.NewTask(task.Title("Task 1"), task.Description("A"), common.GenerateID())
	t1.SetAssignee(assigneeID)

	t2 := task.NewTask(task.Title("Task 2"), task.Description("B"), common.GenerateID())
	t2.SetAssignee(assigneeID)

	t3 := task.NewTask(task.Title("Task 3"), task.Description("C"), common.GenerateID())
	t3.SetAssignee(otherAssignee)

	_ = repo.Save(ctx, t1)
	_ = repo.Save(ctx, t2)
	_ = repo.Save(ctx, t3)

	list, err := repo.GetByAssigneeID(ctx, assigneeID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(list) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(list))
	}
}
