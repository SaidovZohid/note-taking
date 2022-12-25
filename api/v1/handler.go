package v1

import (
	"log"
	"strconv"
	"time"

	"github.com/SaidovZohid/note-taking/api/models"
	"github.com/SaidovZohid/note-taking/config"
	emailPkg "github.com/SaidovZohid/note-taking/pkg/email"
	"github.com/SaidovZohid/note-taking/pkg/utils"
	"github.com/SaidovZohid/note-taking/storage"
	"github.com/SaidovZohid/note-taking/storage/repo"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	cfg      *config.Config
	storage  storage.StorageI
	inMemory storage.InMemoryStorageI
}

type HandlerV1 struct {
	Cfg      *config.Config
	Storage  storage.StorageI
	InMemory storage.InMemoryStorageI
}

func New(options *HandlerV1) *handlerV1 {
	return &handlerV1{
		cfg:      options.Cfg,
		storage:  options.Storage,
		inMemory: options.InMemory,
	}
}

func errResponse(err error) models.ResponseError {
	return models.ResponseError{
		Error: err.Error(),
	}
}

func parseUserModel(user *repo.User) models.GetUserResponse {
	return models.GetUserResponse{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Username:    user.Username,
		ImageUrl:    user.ImageUrl,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   user.UpdatedAt.Format(time.RFC3339),
	}
}

func parseNoteModel(note *repo.Note) models.GetNoteResponse {
	return models.GetNoteResponse{
		ID:          note.ID,
		UserID:      note.UserID,
		Title:       note.Title,
		Description: note.Description,
		CreatedAt:   note.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   note.UpdatedAt.Format(time.RFC3339),
	}
}

func validateUser(ctx *gin.Context) (*models.GetAllUsersParams, error) {
	var (
		limit  int64  = 10
		page   int64  = 1
		sortby string = "desc"
		err    error
	)
	if ctx.Query("limit") != "" {
		limit, err = strconv.ParseInt(ctx.Query("limit"), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	if ctx.Query("page") != "" {
		page, err = strconv.ParseInt(ctx.Query("page"), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	if ctx.Query("sort_by") != "" &&
		(ctx.Query("sort_by") == "desc" || ctx.Query("sort_by") == "asc") {
		sortby = ctx.Query("sort_by")
	}
	return &models.GetAllUsersParams{
		Limit:  limit,
		Page:   page,
		Search: ctx.Query("search"),
		SortBy: sortby,
	}, nil
}

func validateNote(ctx *gin.Context) (*models.GetAllNotesParams, error) {
	var (
		limit  int64  = 10
		page   int64  = 1
		sortby string = "desc"
		err    error
	)
	if ctx.Query("limit") != "" {
		limit, err = strconv.ParseInt(ctx.Query("limit"), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	if ctx.Query("page") != "" {
		page, err = strconv.ParseInt(ctx.Query("page"), 10, 64)
		if err != nil {
			return nil, err
		}
	}
	if ctx.Query("sort_by") != "" &&
		(ctx.Query("sort_by") == "desc" || ctx.Query("sort_by") == "asc") {
		sortby = ctx.Query("sort_by")
	}
	return &models.GetAllNotesParams{
		Limit:  limit,
		Page:   page,
		Search: ctx.Query("search"),
		SortBy: sortby,
	}, nil
}

func (h *handlerV1) sendVerificationCode(key, email string) error {
	code, err := utils.GenerateRandomCode(6)
	if err != nil {
		return err
	}

	err = h.inMemory.Set(key+email, code, time.Minute*10)
	if err != nil {
		return err
	}

	err = emailPkg.SendEmail(h.cfg, &emailPkg.SendEmailRequest{
		To:      []string{email},
		Subject: "Verification Email",
		Body: map[string]string{
			"code": code,
		},
		Type: emailPkg.VerificationEmail,
	})
	if err != nil {
		log.Println("failed to sent code to email")
	}

	return nil
}
