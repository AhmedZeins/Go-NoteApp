package main

import (
	"log"
	"net/http"
	"notes-app/internal/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	notesHandler := handlers.NewNotesHandler()

	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	
	// Serve index.html for root path
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// API Routes
	r.HandleFunc("/api/notes", notesHandler.CreateNote).Methods("POST")
	r.HandleFunc("/api/notes", notesHandler.GetNotes).Methods("GET")
	r.HandleFunc("/api/notes/{id}", notesHandler.GetNote).Methods("GET")
	r.HandleFunc("/api/notes/{id}", notesHandler.UpdateNote).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", notesHandler.DeleteNote).Methods("DELETE")

	// CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Start server
	handler := c.Handler(r)
	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
