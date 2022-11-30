package models

import "time"

type CreateNoteRequest struct {
	Title       string `json:"title" binding:"required,min=2,max=100"`
	Description string `json:"description" binding:"required,max=100"`
}

type GetNoteResponse struct {
	ID          int64      `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type GetAllNotes struct {
	Notes []*GetNoteResponse `json:"notes"`
	Count int64              `json:"count"`
}

type GetAllNotesParams struct {
	Limit  int64  `json:"limit" default:"10"`
	Page   int64  `json:"page" default:"1"`
	Search string `json:"search"`
	SortBy string `json:"sort_by" default:"desc" binding:"oneof=desc asc"`
}
