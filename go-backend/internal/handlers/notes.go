package handlers

import (
	"encoding/json"
	"net/http"
	"notes-app/internal/models"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// NotesHandler handles all note-related operations
type NotesHandler struct {
	notes map[string]models.Note
	mutex sync.RWMutex
}

// NewNotesHandler creates a new NotesHandler
func NewNotesHandler() *NotesHandler {
	return &NotesHandler{
		notes: make(map[string]models.Note),
	}
}

// CreateNote handles the creation of a new note
func (h *NotesHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note.ID = uuid.New().String()
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()

	h.mutex.Lock()
	h.notes[note.ID] = note
	h.mutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// GetNotes returns all notes
func (h *NotesHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	h.mutex.RLock()
	notes := make([]models.Note, 0, len(h.notes))
	for _, note := range h.notes {
		notes = append(notes, note)
	}
	h.mutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

// GetNote returns a specific note by ID
func (h *NotesHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	h.mutex.RLock()
	note, exists := h.notes[id]
	h.mutex.RUnlock()

	if !exists {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}

// UpdateNote updates an existing note
func (h *NotesHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updatedNote models.Note
	if err := json.NewDecoder(r.Body).Decode(&updatedNote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.mutex.Lock()
	if note, exists := h.notes[id]; exists {
		note.Title = updatedNote.Title
		note.Content = updatedNote.Content
		note.UpdatedAt = time.Now()
		h.notes[id] = note
		h.mutex.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(note)
		return
	}
	h.mutex.Unlock()

	http.Error(w, "Note not found", http.StatusNotFound)
}

// DeleteNote deletes a note by ID
func (h *NotesHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	h.mutex.Lock()
	if _, exists := h.notes[id]; exists {
		delete(h.notes, id)
		h.mutex.Unlock()
		w.WriteHeader(http.StatusNoContent)
		return
	}
	h.mutex.Unlock()

	http.Error(w, "Note not found", http.StatusNotFound)
}
