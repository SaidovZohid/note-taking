package v1

import (
	"strconv"

	"github.com/SaidovZohid/note-taking/api/models"
	"github.com/SaidovZohid/note-taking/config"
	"github.com/SaidovZohid/note-taking/storage"
	"github.com/SaidovZohid/note-taking/storage/repo"
	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	cfg     *config.Config
	storage storage.StorageI
}

type HandlerV1 struct {
	Cfg     *config.Config
	Storage *storage.StorageI
}

func New(options *HandlerV1) *handlerV1 {
	return &handlerV1{
		cfg:     options.Cfg,
		storage: *options.Storage,
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
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func parseNoteModel(note *repo.Note) models.GetNoteResponse {
	return models.GetNoteResponse{
		ID:          note.ID,
		UserID:      note.UserID,
		Title:       note.Title,
		Description: note.Description,
		CreatedAt:   note.CreatedAt,
		UpdatedAt:   note.UpdatedAt,
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
