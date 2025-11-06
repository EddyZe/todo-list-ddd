package task

import (
	"log"
	"net/http"
	"todo/internal/adaptar/http/handlers/errcode"
	"todo/internal/adaptar/http/resp"
	"todo/internal/application/task"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *Handler) addComment(c *gin.Context) {
	taskID := c.Param("id")

	body := new(Comment)

	if err := c.ShouldBindJSON(body); err != nil {
		r := resp.ErrorResponse(errcode.CodeInvalidData, "invalid data")
		c.JSON(http.StatusNotFound, r)
		return
	}

	currentUserID := sessions.Default(c).Get("userID").(string)
	log.Println("currentUserID: ", currentUserID)

	if err := h.AddComment(c.Request.Context(), taskID, task.InputAddComment{
		Content:  body.Content,
		AuthorID: currentUserID,
	}); err != nil {
		log.Println(err)
		r := resp.ErrorResponse(errcode.CodeBadRequest, "bad request")
		c.JSON(http.StatusBadRequest, r)
		return
	}

	r := resp.SuccessResponse("")
	c.JSON(http.StatusOK, r)
}
