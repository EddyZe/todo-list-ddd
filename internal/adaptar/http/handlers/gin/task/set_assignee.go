package task

import (
	"net/http"
	"todo/internal/adaptar/http/handlers/errcode"
	"todo/internal/adaptar/http/resp"
	"todo/internal/application/task"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) takeTask(c *gin.Context) {
	currentUserID := sessions.Default(c).Get("userID").(string)
	taskID := c.Param("id")

	if err := h.UpdateAssignee(c.Request.Context(), task.InputUpdateAssignee{
		TaskID:     taskID,
		AssigneeID: currentUserID,
	}); err != nil {
		r := resp.ErrorResponse(errcode.CodeInvalidData, "Не удалось взять задачу")
		c.JSON(http.StatusBadRequest, r)
		return
	}

	r := resp.SuccessResponse("ok")
	c.JSON(http.StatusOK, r)
}
