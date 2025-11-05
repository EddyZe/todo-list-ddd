package sessionauth

import (
	"context"
	"todo/internal/application/authmanager"
	"todo/internal/application/user"

	"github.com/gin-gonic/gin"
)

type Reader interface {
	Authenticate(ctx context.Context, email, password string) (*authmanager.UserOutput, error)
}

type Writer interface {
	CreateUser(ctx context.Context, u user.InputUser) (*user.OutputUser, error)
}

type AuthSessionHandler struct {
	authManager Reader
	userWriter  Writer
}

func New(authManger Reader, userWriter Writer) *AuthSessionHandler {
	return &AuthSessionHandler{
		authManager: authManger,
		userWriter:  userWriter,
	}
}

func (h *AuthSessionHandler) Register(g *gin.RouterGroup) {
	auth := g.Group("/auth")

	auth.POST("/sing-up", h.register)
	auth.POST("/login", h.login)
}
