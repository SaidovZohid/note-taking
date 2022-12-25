package v1

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/SaidovZohid/note-taking/api/models"
	"github.com/SaidovZohid/note-taking/pkg/utils"
	"github.com/SaidovZohid/note-taking/storage/repo"
	"github.com/gin-gonic/gin"
)

var (
	ErrEmailExists          = errors.New("email is already exists")
	ErrCodeExpired          = errors.New("verification code is expired")
	ErrIncorrectCode        = errors.New("verification code is incorrect")
	ErrWrongEmailOrPassword = errors.New("email or password is wrong")
)

const (
	RegisterCodeKey = "register_code_"
)

// @Router /auth/register [post]
// @Summary Create user with token key and get token key.
// @Description Create user with token key and get token key.
// @Tags register
// @Accept json
// @Produce json
// @Param data body models.RegisterRequest true "Data"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) Register(ctx *gin.Context) {
	var (
		req models.RegisterRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	_, err := h.storage.User().GetByEmail(req.Email)
	if !errors.Is(err, sql.ErrNoRows) {
		ctx.JSON(http.StatusBadRequest, errResponse(ErrEmailExists))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user := repo.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hashedPassword,
	}

	userData, err := json.Marshal(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	err = h.inMemory.Set("user_"+req.Email, string(userData), time.Minute*10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	go func() {
		err := h.sendVerificationCode(RegisterCodeKey, req.Email)
		if err != nil {
			log.Printf("failed to send verification code: %v", err)
		}
	}()

	ctx.JSON(http.StatusCreated, models.ResponseOK{
		Message: "Verification code has been sent!",
	})
}

// @Router /auth/verify [post]
// @Summary Verify an user with verification code by email 
// @Description Verify an user with verification code by email 
// @Tags register
// @Accept json
// @Produce json
// @Param data body models.VerifyRequest true "Data"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) Verify(ctx *gin.Context) {
	var (
		req models.VerifyRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	userData, err := h.inMemory.Get("user_" + req.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	var user repo.User
	err = json.Unmarshal([]byte(userData), &user)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errResponse(err))
		return
	}

	code, err := h.inMemory.Get(RegisterCodeKey + user.Email)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errResponse(err))
		return
	}

	if code != req.Code {
		ctx.JSON(http.StatusForbidden, errResponse(ErrIncorrectCode))
		return
	}

	result, err := h.storage.User().Create(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	token, _, err := utils.CreateToken(h.cfg, &utils.TokenParams{
		UserID:   user.ID,
		Email:    user.Email,
		Duration: time.Minute * 10,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.AuthResponse{
		ID:          result.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		AccessToken: token,
	})
}

// @Router /auth/login [post]
// @Summary log in with email password after verifing user
// @Description log in with email password after verifing user
// @Tags register
// @Accept json
// @Produce json
// @Param data body models.LoginRequest true "Data"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) Login(ctx *gin.Context) {
	var (
		req models.LoginRequest
	)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := h.storage.User().GetByEmail(req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusForbidden, errResponse(ErrWrongEmailOrPassword))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errResponse(ErrWrongEmailOrPassword))
		return
	}

	token, _, err := utils.CreateToken(h.cfg, &utils.TokenParams{
		UserID:   user.ID,
		Email:    user.Email,
		Duration: time.Hour * 24 * 30,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.AuthResponse{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		AccessToken: token,
	})
}
