package user

import (
	"errors"
	"net/http"
	"todo/internal/adaptar/http/handlers/errcode"
	"todo/internal/adaptar/http/resp"
	"todo/internal/application/user"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getID(c *gin.Context) {
	userID := c.Param("id")

	usr, err := h.GetByID(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			r := resp.ErrorResponse(errcode.CodeNotFound, "Пользователь с таким ID не найден")
			c.JSON(http.StatusNotFound, r)
			return
		}

		r := resp.ErrorResponse(errcode.CodeInternalError, "ошибка на стороне сервера")
		c.JSON(http.StatusInternalServerError, r)
		return
	}

	r := resp.SuccessResponse(User{
		ID:    usr.ID,
		Email: usr.Email,
	})

	c.JSON(http.StatusOK, r)
}
