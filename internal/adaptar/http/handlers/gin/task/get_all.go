package task

import (
	"net/http"
	"todo/internal/adaptar/http/resp"
	"todo/internal/application/task"

	"github.com/gin-gonic/gin"
)

func (h *Handler) All(c *gin.Context) {
	assigneeID := c.Query("assignee-id")
	authorID := c.Query("author-id")

	tasks := make([]*task.OutputTask, 0)

	if assigneeID != "" {
		t := h.GetByAssigneeID(c.Request.Context(), assigneeID)
		tasks = append(tasks, t...)
	}

	if authorID != "" {
		t := h.GetByAuthorID(c.Request.Context(), authorID)
		tasks = append(tasks, t...)
	}

	if assigneeID == "" && authorID == "" {
		t := h.GetAll(c.Request.Context())
		tasks = append(tasks, t...)
	}

	res := make([]*OutputTask, 0, len(tasks))

	for _, t := range tasks {
		res = append(res, MapTaskToHandler(t))
	}

	r := resp.SuccessResponse(ListResponse{
		Tasks: res,
	})

	c.JSON(http.StatusOK, r)
}
