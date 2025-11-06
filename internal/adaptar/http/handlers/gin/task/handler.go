package task

import (
	"context"
	"todo/internal/application/task"
	"todo/internal/infrastructure/middleware"

	"github.com/gin-gonic/gin"
)

type Writer interface {
	CreateTask(ctx context.Context, input task.InputCreateTask) (output *task.OutputTask, err error)
	DeleteByID(ctx context.Context, taskID string) error
	DeleteComment(ctx context.Context, comment task.InputRemoveComment) error
	UpdateAssignee(ctx context.Context, input task.InputUpdateAssignee) error
	AddComment(ctx context.Context, taskID string, c task.InputAddComment) error
	ChangeStatus(ctx context.Context, taskID string, status string) error
}

type Reader interface {
	GetByID(ctx context.Context, taskID string) (*task.OutputTask, error)
	GetByAuthorID(ctx context.Context, authorID string) []*task.OutputTask
	GetByAssigneeID(ctx context.Context, assigneeID string) []*task.OutputTask
	GetAll(ctx context.Context) []*task.OutputTask
}

type UseCase interface {
	Reader
	Writer
}

type Handler struct {
	UseCase
}

func NewHandler(useCase UseCase) *Handler {
	return &Handler{
		useCase,
	}
}

func (h *Handler) Register(c *gin.RouterGroup) {
	tasks := c.Group("/tasks")
	tasks.Use(middleware.AuthRequiredGin())
	{
		tasks.GET("", h.All)
		tasks.POST("", h.createTask)
		tasks.GET("/:id", h.byID)
		tasks.DELETE("/:id", h.delete)
		tasks.POST("/:id/addComment", h.addComment)
		tasks.PATCH("/:id/changeStatus", h.changeStatus)
		tasks.PATCH("/:id/takeTask", h.takeTask)
	}
}
