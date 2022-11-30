package postgres_test

import (
	"testing"

	"github.com/SaidovZohid/note-taking/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createNote(t *testing.T) *repo.Note {
	note, err := dbManager.Note().Create(&repo.Note{
		Title: faker.FirstName(),
		Description: faker.Sentence(),
		UserID: 26,
	})
	require.NoError(t, err)
	return note
}

func deleteNote(t *testing.T, note_id int64) {
	err := dbManager.Note().Delete(note_id)
	require.NoError(t, err)
}

func TestCreateNote(t *testing.T) {
	note := createNote(t)
	deleteNote(t, note.ID)
	require.NotEmpty(t, note)
}

func TestGetNote(t *testing.T) {
	note := createNote(t)
	require.NotEmpty(t, note)
	note2, err := dbManager.Note().Get(note.ID)
	deleteNote(t, note.ID)
	require.NoError(t, err)
	require.NotEmpty(t, note2)
}

func TestUpdateNote(t *testing.T) {
	note := createNote(t)
	note2, err := dbManager.Note().Update(&repo.Note{
		UserID: 26,
		Title: faker.ChineseFirstName(),
		Description: faker.Sentence(),
	})
	deleteNote(t, note.ID)
	require.NoError(t, err)
	require.NotEmpty(t, note)
	require.NotEmpty(t, note2)
}

func TestDeleteNote(t *testing.T) {
	note := createNote(t)
	deleteNote(t, note.ID)
}

func TestGetAllNotes(t *testing.T) {
	note := createNote(t)
	notes, err := dbManager.Note().GetAll(&repo.GetAllNotesParams{
		Limit: 10,
		Page: 1,
		SortBy: "ASC",
	})
	deleteNote(t, note.ID)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(notes.Notes), 1)
	require.NotEmpty(t, note)
}