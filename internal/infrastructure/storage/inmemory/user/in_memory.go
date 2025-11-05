package user

import (
	"context"
	"errors"
	"fmt"
	"math"
	"todo/internal/domain/common"
	"todo/internal/domain/user"
)

// InMemory реализует user.Repository
type InMemory struct {
	memory map[common.ID]*user.User
}

// NewInMemory создаёт новый репозиторий в памяти
func NewInMemory() *InMemory {
	return &InMemory{
		memory: make(map[common.ID]*user.User),
	}
}

// ---------------- Writer ----------------

func (r *InMemory) Save(_ context.Context, u *user.User) error {
	if u == nil {
		return fmt.Errorf("%w: пользователь не может быть nil", user.ErrInvalidUser)
	}

	if ok := r.ExistsByEmail(context.Background(), u.Email()); ok {
		return fmt.Errorf("%w: пользователь с таким email уже существует", user.ErrUserAlreadyExists)
	}

	r.memory[u.ID()] = u
	return nil
}

func (r *InMemory) DeleteByID(_ context.Context, id common.ID) error {
	if _, ok := r.memory[id]; !ok {
		return fmt.Errorf("%w: пользователь не найден", user.ErrUserNotFound)
	}
	delete(r.memory, id)
	return nil
}

// ---------------- Reader ----------------

func (r *InMemory) GetByID(_ context.Context, id common.ID) (*user.User, error) {
	u, ok := r.memory[id]
	if !ok {
		return nil, fmt.Errorf("%w: пользователь не найден", user.ErrUserNotFound)
	}
	return u, nil
}

func (r *InMemory) GetByEmail(_ context.Context, email user.Email) (*user.User, error) {
	for _, u := range r.memory {
		if u.Email().String() == email.String() {
			return u, nil
		}
	}
	return nil, fmt.Errorf("%w: пользователь не найден", user.ErrUserNotFound)
}

func (r *InMemory) GetAll(_ context.Context, page, limit int) ([]*user.User, int, error) {
	if page < 1 {
		return make([]*user.User, 0), 0, errors.New("страница должна быть больше чем 0")
	}

	if limit <= 0 {
		return nil, 0, errors.New("limit должен быть больше 0")
	}

	all := make([]*user.User, 0, len(r.memory))
	for _, u := range r.memory {
		all = append(all, u)
	}

	totalPages := int(math.Ceil(float64(len(all)) / float64(limit)))
	offset := (page - 1) * limit

	// простая пагинация
	if offset >= len(all) {
		return []*user.User{}, totalPages, nil
	}

	end := offset + limit
	if end > len(all) {
		end = len(all)
	}

	return all[offset:end], totalPages, nil
}

func (r *InMemory) ExistsByEmail(_ context.Context, email user.Email) bool {
	for _, u := range r.memory {
		if u.Email().String() == email.String() {
			return true
		}
	}
	return false
}
