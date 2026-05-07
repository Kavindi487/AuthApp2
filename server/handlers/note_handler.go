package handlers

import (
    "server/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type NoteHandler struct {
    NoteService *services.NoteService
}

func NewNoteHandler(noteService *services.NoteService) *NoteHandler {
    return &NoteHandler{NoteService: noteService}
}

type CreateNoteRequest struct {
    Title   string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
}

func (h *NoteHandler) CreateNote(c *gin.Context) {
    // Get user_id set by JWT middleware
    userID := c.MustGet("user_id").(uint)

    var req CreateNoteRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    note, err := h.NoteService.CreateNote(userID, req.Title, req.Content)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"note": note})
}

func (h *NoteHandler) GetMyNotes(c *gin.Context) {
    userID := c.MustGet("user_id").(uint)

    notes, err := h.NoteService.GetMyNotes(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch notes"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"notes": notes})
}

func (h *NoteHandler) DeleteNote(c *gin.Context) {
    userID := c.MustGet("user_id").(uint)

    // Read :id from the URL
    noteIDStr := c.Param("id")
    noteID, err := strconv.Atoi(noteIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid note id"})
        return
    }

    err = h.NoteService.DeleteNote(uint(noteID), userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "could not delete note"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "note deleted"})
}