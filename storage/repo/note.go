package repo

import "time"

type Note struct {
	ID          int64
	UserID      int64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type NoteStorageI interface {
	Create(n *Note) (*Note, error)
	Get(note_id int64) (*Note, error)
	Update(n *Note) (*Note, error)
	Delete(noteId, userId int64) error
	GetAll(params *GetAllNotesParams) (*GetALlNotesResult, error)
}

type GetAllNotesParams struct {
	Limit  int64
	Page   int64
	Search string
	SortBy string
}

type GetALlNotesResult struct {
	Notes []*Note
	Count int64
}
