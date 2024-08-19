package handlers

import (
	"ShareGo/api/db"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	userIDStr := r.FormValue("user_id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := os.MkdirAll("./uploads", os.ModePerm); err != nil {
		http.Error(w, "Error creating upload directory", http.StatusInternalServerError)
		return
	}

	filePath := "./uploads/" + time.Now().Format("20060102150405") + "-" + fileHeader.Filename
	out, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error creating the file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		http.Error(w, "Error saving the file", http.StatusInternalServerError)
		return
	}

	fileInfo, err := db.CreateFile(fileHeader.Filename, filePath, userID)
	if err != nil {
		http.Error(w, "Error saving file info to database", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"success":  true,
		"filename": fileInfo.Filename,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response as JSON", http.StatusInternalServerError)
		return
	}
}
