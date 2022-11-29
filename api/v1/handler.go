package v1

import (
	"github.com/SaidovZohid/note-taking/api/models"
	"github.com/SaidovZohid/note-taking/config"
	"github.com/SaidovZohid/note-taking/storage"
	"github.com/SaidovZohid/note-taking/storage/repo"
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

func ParseModel(user *repo.User) models.GetUserResponse {
	return models.GetUserResponse{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.LastName,
		PhoneNumber: user.PhoneNumber,
		Username:    user.Username,
		ImageUrl:    user.ImageUrl,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}
