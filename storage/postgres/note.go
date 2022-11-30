package postgres

import (
	"time"

	"github.com/SaidovZohid/note-taking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type noteRepo struct {
	db *sqlx.DB
}

func NewNote(db *sqlx.DB) repo.NoteStorageI {
	return &noteRepo{db: db}
}

func (nr *noteRepo) Create(n *repo.Note) (*repo.Note, error) {
	query := `
		INSERT INTO notes (
			user_id,
			title,
			description
		) VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at 
	`
	err := nr.db.QueryRow(
		query,
		n.UserID,
		n.Title,
		n.Description,
	).Scan(
		&n.ID,
		&n.CreatedAt,
		&n.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	
	return n, nil
}

func (nr *noteRepo) Get(note_id int64) (*repo.Note, error) {
	var (
		note repo.Note
	)
	query := `
		SELECT 
			id,
			user_id, 
			title,
			description,
			created_at,
			updated_at
		FROM notes WHERE id = $1
	`
	err := nr.db.QueryRow(query, note_id).Scan(
		&note.ID,
		&note.UserID,
		&note.Title,
		&note.Description,
		&note.CreatedAt,
		&note.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (nr *noteRepo) Update(n *repo.Note) (*repo.Note, error) {
	var (
		note repo.Note
	)

	query := `
		UPDATE notes SET 
			title = $1,
			description = $2,
			updated_at = $3
		RETURNING 
			id,
			user_id, 
			title,
			description,
			created_at,
			updated_at
	`

	err := nr.db.QueryRow(
		query,
		n.Title,
		n.Description,
		time.Now(),
	).Scan(
		&note.ID,
		&note.UserID,
		&note.Title,
		&note.Description,
		&note.CreatedAt,
		&note.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &note, err
}

func (nr *noteRepo) Delete(note_id int64) error {
	query := `
		UPDATE notes SET 
			deleted_at = $1
		WHERE id = $2
	`

	_, err := nr.db.Exec(query, time.Now())
	if err != nil {
		return err
	}

	return nil
}