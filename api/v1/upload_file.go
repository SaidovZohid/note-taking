package v1

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/SaidovZohid/note-taking/api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

// @Security ApiKeyAuth
// @Router /file-upload [post]
// @Summary File upload
// @Description File upload
// @Tags file-upload
// @Accept json
// @Produce json
// @Param file formData file true "File"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UploadFile(ctx *gin.Context) {
	var file File

	if err := ctx.ShouldBind(&file); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	id := uuid.New()
	fileName := id.String() + filepath.Ext(file.File.Filename)
	dir, _ := os.Getwd()

	if _, err := os.Stat(dir + "/media"); os.IsNotExist(err) {
		err = os.Mkdir(dir + "/media", os.ModePerm)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return
		}
	} 	

	filePath := "/media/" + fileName
	err := ctx.SaveUploadedFile(file.File, dir + filePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, models.ResponseOK{
		Message: filePath,
	})
}