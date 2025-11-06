package user

import (
	"net/http"
	"strconv"
	"todo/internal/adaptar/http/handlers/errcode"
	"todo/internal/adaptar/http/resp"
	"todo/internal/application/user"

	"github.com/gin-gonic/gin"
)

func (h *Handler) all(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		r := resp.ErrorResponse(errcode.CodeBadRequest, "page - должно быть числом")
		c.JSON(http.StatusBadRequest, r)
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		r := resp.ErrorResponse(errcode.CodeBadRequest, "limit - должен быть числом")
		c.JSON(http.StatusBadRequest, r)
		return
	}

	users := h.UseCase.GetAll(c.Request.Context(), user.InputUserList{
		Page:  page,
		Limit: limit,
	})

	respUsers := make([]User, 0, len(users.Users))

	for _, u := range users.Users {
		respUsers = append(respUsers, User{
			ID:    u.ID,
			Email: u.Email,
		})
	}

	r := resp.SuccessResponse(AllUserResp{
		Users:       respUsers,
		TotalPages:  users.TotalPages,
		CurrentPage: users.CurrentPage,
		Limit:       users.Limit,
	})

	c.JSON(http.StatusOK, r)
}
