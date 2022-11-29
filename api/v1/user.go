package v1

import (
	"net/http"

	"github.com/SaidovZohid/note-taking/api/models"
	"github.com/SaidovZohid/note-taking/pkg/utils"
	"github.com/SaidovZohid/note-taking/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Router /users [post]
// @Summary Create a user
// @Description Create a user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User"
// @Success 201 {object} models.GetUserResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateUser(ctx *gin.Context) {
	var (
		req models.CreateUserRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	user, err := h.storage.User().Create(&repo.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		PhoneNumber: req.PhoneNumber,
		Email: req.Email,
		ImageUrl: req.ImageUrl,
		Username: req.Username,
		Password: hashedPassword,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, ParseModel(user))
}