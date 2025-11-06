package task

import (
	"net/http"
	"todo/internal/adaptar/http/handlers/errcode"
	"todo/internal/adaptar/http/resp"

	"github.com/gin-gonic/gin"
)

func (h *Handler) byID(c *gin.Context) {
	id := c.Param("id")

	t, err := h.GetByID(c.Request.Context(), id)
	if err != nil {
		r := resp.ErrorResponse(errcode.CodeNotFound, "Задача с таким ID не найдена")
		c.JSON(http.StatusNotFound, r)
		return
	}

	output := MapTaskToHandler(t)
	r := resp.SuccessResponse(output)
	c.JSON(http.StatusOK, r)
}
