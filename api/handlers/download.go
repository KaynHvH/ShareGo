package handlers

import (
	"ShareGo/api/db"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
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

	file, err := os.Open(fileRecord.Filepath)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+fileRecord.Filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, file)
}
