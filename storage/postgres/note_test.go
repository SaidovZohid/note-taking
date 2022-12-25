package postgres_test

import (
	"testing"

	"github.com/SaidovZohid/note-taking/storage/repo"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createNote(t *testing.T) *repo.Note {
	user := createUser(t)
	note, err := dbManager.Note().Create(&repo.Note{
		Title:       faker.FirstName(),
		Description: faker.Sentence(),
		UserID:      user.ID,
	})
	require.NoError(t, err)
	return note
}

func deleteNote(t *testing.T, noteId, userid int64) {
	err := dbManager.Note().Delete(noteId, userid)
	require.NoError(t, err)
}

func TestCreateNote(t *testing.T) {
	note := createNote(t)
	require.NotEmpty(t, note)
	deleteNote(t, note.ID, note.UserID)
}

func TestGetNote(t *testing.T) {
	note := createNote(t)
	require.NotEmpty(t, note)
	note2, err := dbManager.Note().Get(note.ID)
	require.NoError(t, err)
	require.NotEmpty(t, note2)
	deleteNote(t, note.ID, note.UserID)
}

func TestUpdateNote(t *testing.T) {
	note := createNote(t)
	note2, err := dbManager.Note().Update(&repo.Note{
		ID:          note.ID,
		UserID:      note.UserID,
		Title:       faker.ChineseFirstName(),
		Description: faker.Sentence(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, note)
	require.NotEmpty(t, note2)
	deleteNote(t, note.ID, note.UserID)
}

func TestDeleteNote(t *testing.T) {
	note := createNote(t)
	deleteNote(t, note.ID, note.UserID)
}

func TestGetAllNotes(t *testing.T) {
	note := createNote(t)
	notes, err := dbManager.Note().GetAll(&repo.GetAllNotesParams{
		Limit:  10,
		Page:   1,
		SortBy: "ASC",
	})
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(notes.Notes), 1)
	require.NotEmpty(t, note)
	deleteNote(t, note.ID, note.UserID)
}
