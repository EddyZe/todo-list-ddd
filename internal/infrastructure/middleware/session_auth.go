package middleware

import (
	"net/http"
	"todo/internal/adaptar/http/handlers/errcode"
	"todo/internal/adaptar/http/resp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequiredGin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("userID")
		if userID == nil {
			r := resp.ErrorResponse(errcode.CodeUnauthorized, "Пройдите авторизацию")
			c.JSON(http.StatusUnauthorized, r)
			c.Abort()
			return
		}

		email := session.Get("email")
		if email == nil {
			r := resp.ErrorResponse(errcode.CodeUnauthorized, "Пройдите авторизацию")
			c.JSON(http.StatusUnauthorized, r)
			c.Abort()
			return
		}

		c.Next()
	}
}
