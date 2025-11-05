package sessionauth

import (
	"errors"
	"net/http"
	"todo/internal/adaptar/http/handlers/errcode"
	"todo/internal/adaptar/http/resp"
	"todo/internal/application/user"

	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponse struct {
	ID string `json:"id"`
}

func (h *AuthSessionHandler) register(c *gin.Context) {
	req := new(createUserRequest)

	if err := c.ShouldBind(req); err != nil {
		r := resp.ErrorResponse(errcode.CodeInvalidData, "Не верное тело запроса")
		c.JSON(http.StatusBadRequest, r)
		return
	}

	inputUser := user.InputUser{
		Email:    req.Email,
		Password: req.Password,
	}

	output, err := h.userWriter.CreateUser(c.Request.Context(), inputUser)
	if err != nil {
		if errors.Is(err, user.ErrUserIsExists) {
			r := resp.ErrorResponse(errcode.CodeIsExists, "Пользователь с таким email существует")
			c.JSON(http.StatusBadRequest, r)
			return
		}

		if errors.Is(err, user.ErrCreatingUser) {
			r := resp.ErrorResponse(errcode.CodeInvalidData, "Проверьте введенные данные")
			c.JSON(http.StatusBadRequest, r)
			return
		}

		r := resp.ErrorResponse(errcode.CodeInternalError, "ошибка сервера")
		c.JSON(http.StatusInternalServerError, r)
		return
	}

	response := createUserResponse{ID: output.ID}
	c.JSON(http.StatusOK, resp.SuccessResponse(response))
}
