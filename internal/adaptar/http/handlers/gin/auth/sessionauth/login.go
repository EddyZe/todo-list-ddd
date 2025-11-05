package sessionauth

import (
	"errors"
	"net/http"
	"todo/internal/adaptar/http/handlers/errcode"
	"todo/internal/adaptar/http/resp"
	"todo/internal/application/authmanager"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type userLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userLoginResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (h *AuthSessionHandler) login(c *gin.Context) {
	req := new(userLoginRequest)

	if err := c.BindJSON(req); err != nil {
		r := resp.ErrorResponse(errcode.CodeInvalidData, "Невалидное тело запроса")
		c.JSON(http.StatusBadRequest, r)
		return
	}

	u, err := h.authManager.Authenticate(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, authmanager.ErrInvalidCredentials) {
			r := resp.ErrorResponse(errcode.CodeInvalidCredentials, "Не верный логин или пароль")
			c.JSON(http.StatusUnauthorized, r)
			return
		}
		r := resp.ErrorResponse(errcode.CodeInternalError, "Произошла ошибка на сервере")
		c.JSON(http.StatusInternalServerError, r)
		return
	}

	r := resp.SuccessResponse(userLoginResponse{
		ID:    u.ID,
		Email: u.Email,
	})

	session := sessions.Default(c)
	session.Set("userID", u.ID)
	session.Set("email", u.Email)

	if err := session.Save(); err != nil {
		r := resp.ErrorResponse(errcode.CodeInternalError, "ошибка при установки сессии")
		c.JSON(http.StatusInternalServerError, r)
		return
	}

	c.JSON(http.StatusOK, r)

	return
}
