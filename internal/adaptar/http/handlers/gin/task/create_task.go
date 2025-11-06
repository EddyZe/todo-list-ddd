package task

import (
	"log"
	"net/http"
	"todo/internal/adaptar/http/handlers/errcode"
	"todo/internal/adaptar/http/resp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(c *gin.Context) {
	req := new(InputCreateTask)
	if err := c.ShouldBind(req); err != nil {
		r := resp.ErrorResponse(errcode.CodeInvalidData, "невалидное тело запроса")
		c.JSON(http.StatusBadRequest, r)
		return
	}

	currentUser := sessions.Default(c).Get("userID").(string)

	reqMap := MapCreateTaskInput(*req, currentUser)

	r, err := h.CreateTask(c.Request.Context(), *reqMap)
	if err != nil {
		log.Println(err)
		rr := resp.ErrorResponse(errcode.CodeInternalError, "ошибка создания задачи")
		c.JSON(http.StatusInternalServerError, rr)
		return
	}

	mapTask := MapTaskToHandler(r)

	rr := resp.SuccessResponse(mapTask)
	c.JSON(http.StatusOK, rr)
}
