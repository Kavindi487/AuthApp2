package repository

import (
    "server/models"
    "gorm.io/gorm"
)

type NoteRepository struct {
    DB *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
    return &NoteRepository{DB: db}
}

func (r *NoteRepository) CreateNote(note *models.Note) error {
    return r.DB.Create(note).Error
}

func (r *NoteRepository) GetNotesByUserID(userID uint) ([]models.Note, error) {
    var notes []models.Note
    result := r.DB.Where("user_id = ?", userID).
                  Order("created_at desc").
                  Find(&notes)
    return notes, result.Error
}

func (r *NoteRepository) DeleteNote(noteID uint, userID uint) error {
    // userID check ensures users can only delete their OWN notes
    return r.DB.Where("id = ? AND user_id = ?", noteID, userID).
               Delete(&models.Note{}).Error
}