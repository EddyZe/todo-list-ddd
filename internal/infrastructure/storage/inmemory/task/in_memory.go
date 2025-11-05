package task

import (
	"context"
	"todo/internal/domain/common"
	"todo/internal/domain/task"
)

// InMemory реализует task.Repository
type InMemory struct {
	memory map[common.ID]*task.Task
}

// NewInMemory создаёт новый репозиторий в памяти
func NewInMemory() *InMemory {
	return &InMemory{
		memory: make(map[common.ID]*task.Task),
	}
}

// ---------------- Writer ----------------

func (r *InMemory) Save(_ context.Context, t *task.Task) error {
	if t == nil {
		return task.ErrInvalidTask
	}
	r.memory[t.ID()] = t
	return nil
}

func (r *InMemory) DeleteByID(_ context.Context, id common.ID) error {
	if _, ok := r.memory[id]; !ok {
		return task.ErrTaskNotFound
	}
	delete(r.memory, id)
	return nil
}

// ---------------- Reader ----------------

func (r *InMemory) GetByID(_ context.Context, id common.ID) (*task.Task, error) {
	t, ok := r.memory[id]
	if !ok {
		return nil, task.ErrTaskNotFound
	}
	return t, nil
}

func (r *InMemory) GetByAuthorID(_ context.Context, authorID common.ID) ([]*task.Task, error) {
	var result []*task.Task
	for _, t := range r.memory {
		if t.CreatorID() == authorID {
			result = append(result, t)
		}
	}
	return result, nil
}

func (r *InMemory) GetByAssigneeID(_ context.Context, assigneeID common.ID) ([]*task.Task, error) {
	var result []*task.Task
	for _, t := range r.memory {
		if t.AssigneeID() == assigneeID {
			result = append(result, t)
		}
	}
	return result, nil
}
