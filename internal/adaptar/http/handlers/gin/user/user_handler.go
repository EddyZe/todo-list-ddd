package user

import (
	"context"
	"todo/internal/application/user"
	"todo/internal/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

type Reader interface {
	GetByID(ctx context.Context, id string) (*user.OutputUser, error)
	GetByEmail(ctx context.Context, email string) (*user.OutputUser, error)
	GetAll(ctx context.Context, input user.InputUserList) *user.OutputUserList
}

type Writer interface {
	CreateUser(ctx context.Context, u user.InputUser) (*user.OutputUser, error)
}

type UseCase interface {
	Reader
	Writer
}

type Handler struct {
	UseCase
}

func New(u UseCase) *Handler {
	return &Handler{
		u,
	}
}

func (h *Handler) Register(g *gin.RouterGroup) {
	users := g.Group("/users")

	users.Use(middleware.AuthRequiredGin())
	{
		users.GET("", h.all)
		users.GET("/:id", h.getID)
	}
}
