package repo

import "time"

type Note struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type NoteStorageI interface {
	Create(n *Note) (*Note, error)
}