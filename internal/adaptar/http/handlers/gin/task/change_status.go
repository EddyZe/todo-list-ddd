package task

import (
	"net/http"
	"todo/internal/adaptar/http/handlers/errcode"
	"todo/internal/adaptar/http/resp"

	"github.com/gin-gonic/gin"
)

func (h *Handler) changeStatus(c *gin.Context) {
	taskID := c.Param("id")

	type statusReq struct {
		Status string `json:"status" binding:"required"`
	}

	var req statusReq
	if err := c.ShouldBindJSON(&req); err != nil {
		r := resp.ErrorResponse(errcode.CodeInvalidData, "невалидное тело запроса")
		c.JSON(http.StatusBadRequest, r)
		return
	}

	if err := h.ChangeStatus(c.Request.Context(), taskID, req.Status); err != nil {
		r := resp.ErrorResponse(errcode.CodeInvalidData, err.Error())
		c.JSON(http.StatusBadRequest, r)
		return
	}

	r := resp.SuccessResponse("")
	c.JSON(http.StatusOK, r)
}
