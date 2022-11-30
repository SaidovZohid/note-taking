package v1

import (
	"net/http"

	"github.com/SaidovZohid/note-taking/api/models"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) CreateNote(ctx *gin.Context) {
	var (
		req models.CreateNoteRequest
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	
}