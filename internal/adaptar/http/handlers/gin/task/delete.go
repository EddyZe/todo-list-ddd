package task

import (
	"net/http"
	"todo/internal/adaptar/http/handlers/errcode"
	"todo/internal/adaptar/http/resp"

	"github.com/gin-gonic/gin"
)

func (h *Handler) delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.DeleteByID(c.Request.Context(), id); err != nil {
		r := resp.ErrorResponse(errcode.CodeNotFound, "Задача не найдена")
		c.JSON(http.StatusNotFound, r)
		return
	}

	r := resp.SuccessResponse("")
	c.JSON(http.StatusOK, r)
}
