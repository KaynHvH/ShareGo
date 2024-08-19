package handlers

import (
	"ShareGo/api/db"
	"encoding/json"
	"net/http"
)

func FilesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := db.GetAllFiles()
	if err != nil {
		http.Error(w, "Error fetching files", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(files); err != nil {
		http.Error(w, "Error encoding files as JSON", http.StatusInternalServerError)
		return
	}
}
