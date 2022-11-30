package storage

import (
	"github.com/SaidovZohid/note-taking/storage/postgres"
	"github.com/SaidovZohid/note-taking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
	Note() repo.NoteStorageI
}

type StoragePg struct {
	userRepo repo.UserStorageI
	noteRepo repo.NoteStorageI
}

func NewStorage(db *sqlx.DB) StorageI {
	return &StoragePg{
		userRepo: postgres.NewUserStorage(db),
		noteRepo: postgres.NewNote(db),
	}
}

func (s *StoragePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s *StoragePg) Note() repo.NoteStorageI {
	return s.noteRepo
}