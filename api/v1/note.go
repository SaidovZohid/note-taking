package v1

import (
	"net/http"
	"strconv"

	"github.com/SaidovZohid/note-taking/api/models"
	"github.com/SaidovZohid/note-taking/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /notes [post]
// @Summary Create a note
// @Description Create a note
// @Tags note
// @Accept json
// @Produce json
// @Param user body models.CreateOrUpdateNoteRequest true "User"
// @Success 200 {object} models.GetNoteResponse
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) CreateNote(ctx *gin.Context) {
	var (
		req models.CreateOrUpdateNoteRequest
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	note, err := h.storage.Note().Create(&repo.Note{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, models.GetNoteResponse{
		ID:          note.ID,
		UserID:      note.UserID,
		Title:       note.Title,
		Description: note.Description,
		CreatedAt:   note.CreatedAt,
		UpdatedAt:   note.UpdatedAt,
	})
}

// @Security ApiKeyAuth
// @Router /notes/{id} [get]
// @Summary get a note by id
// @Description get a note by id
// @Tags note
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.GetNoteResponse
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) GetNote(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	note, err := h.storage.Note().Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.GetNoteResponse{
		ID:          note.ID,
		UserID:      note.UserID,
		Title:       note.Title,
		Description: note.Description,
		CreatedAt:   note.CreatedAt,
		UpdatedAt:   note.UpdatedAt,
	})
}

// @Security ApiKeyAuth
// @Router /notes/{id} [put]
// @Summary Update a note
// @Description Update a note
// @Tags note
// @Accept json
// @Produce json
// @Param user body models.CreateOrUpdateNoteRequest true "User"
// @Success 200 {object} models.GetNoteResponse
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) UpdateNote(ctx *gin.Context) {
	var (
		req models.CreateOrUpdateNoteRequest
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	note, err := h.storage.Note().Update(&repo.Note{
		ID:          req.UserID,
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.GetNoteResponse{
		ID:          note.ID,
		UserID:      note.UserID,
		Title:       note.Title,
		Description: note.Description,
		CreatedAt:   note.CreatedAt,
		UpdatedAt:   note.UpdatedAt,
	})
}

// @Security ApiKeyAuth
// @Router /notes/{id} [delete]
// @Summary delete a note
// @Description delete a note
// @Tags note
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) DeleteNote(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err = h.storage.Note().Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseOK{
		Message: "Succesfully deleted!",
	})
}

// @Router /notes [get]
// @Summary get all note
// @Description get all note
// @Tags note
// @Accept json
// @Produce json
// @Param param query models.GetAllNotesParams false "Param"
// @Success 200 {object} models.GetAllNotes
// @Failure 500 {object} models.ResponseError
// @Failure 400 {object} models.ResponseError
func (h *handlerV1) GetAllNotes(ctx *gin.Context) {
	params, err := validateNote(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	notes, err := h.storage.Note().GetAll(&repo.GetAllNotesParams{
		Limit:  params.Limit,
		Page:   params.Page,
		Search: params.Search,
		SortBy: params.SortBy,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, getresponse(notes))
}

func getresponse(notes *repo.GetALlNotesResult) *models.GetAllNotes {
	var res models.GetAllNotes
	for _, note := range notes.Notes {
		n := parseNoteModel(note)
		res.Notes = append(res.Notes, &n)
	}
	res.Count = notes.Count
	return &res
}
