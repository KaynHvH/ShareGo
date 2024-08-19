package handlers

import (
	"ShareGo/api/db"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	fileID := mux.Vars(r)["id"]
	id, err := strconv.Atoi(fileID)
	if err != nil {
		http.Error(w, "Invalid file ID", http.StatusBadRequest)
		return
	}

	fileRecord, err := db.GetFileByID(id)
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusInternalServerError)
		return
	}
	if fileRecord == nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	if err := os.Remove(fileRecord.Filepath); err != nil {
		http.Error(w, "Error deleting file", http.StatusInternalServerError)
		return
	}

	if err := db.DeleteFile(id); err != nil {
		http.Error(w, "Error deleting file record", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
