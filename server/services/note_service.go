package services

import (
    "server/models"
    "server/repository"
    "errors"
)

type NoteService struct {
    NoteRepo *repository.NoteRepository
}

func NewNoteService(noteRepo *repository.NoteRepository) *NoteService {
    return &NoteService{NoteRepo: noteRepo}
}

func (s *NoteService) CreateNote(userID uint, title, content string) (*models.Note, error) {
    if title == "" || content == "" {
        return nil, errors.New("title and content are required")
    }

    note := &models.Note{
        UserID:  userID,
        Title:   title,
        Content: content,
    }

    err := s.NoteRepo.CreateNote(note)
    if err != nil {
        return nil, err
    }
    return note, nil
}

func (s *NoteService) GetMyNotes(userID uint) ([]models.Note, error) {
    return s.NoteRepo.GetNotesByUserID(userID)
}

func (s *NoteService) DeleteNote(noteID uint, userID uint) error {
    return s.NoteRepo.DeleteNote(noteID, userID)
}