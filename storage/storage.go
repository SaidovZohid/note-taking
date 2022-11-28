package storage

import (
	"github.com/SaidovZohid/note-taking/storage/postgres"
	"github.com/SaidovZohid/note-taking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
}

type StoragePg struct {
	userRepo repo.UserStorageI
}

func New(db *sqlx.DB) StorageI {
	return &StoragePg{
		userRepo: postgres.NewUserStorage(db),
	}
}

func (s *StoragePg) User() repo.UserStorageI {
	return s.userRepo
}