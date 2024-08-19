package routes

import (
	"ShareGo/api/auth"
	"ShareGo/api/handlers"
	"ShareGo/api/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func InitRoutes(router *mux.Router) {
	// Middlewares
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.AuthenticationMiddleware)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// Html sites
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/register.html")
	})

	router.HandleFunc("/register.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/register.html")
	})

	router.HandleFunc("/login.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/login.html")
	})

	router.HandleFunc("/uploadfiles", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	router.HandleFunc("/files.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/files.html")
	})

	// Users endpoints
	router.HandleFunc("/register", auth.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", auth.LoginHandler).Methods("POST")

	// Files endpoints
	router.HandleFunc("/files/{id:[0-9]+}/download", handlers.DownloadHandler).Methods("GET")
	router.HandleFunc("/files/{id:[0-9]+}", handlers.DeleteHandler).Methods("DELETE")
	router.HandleFunc("/files", handlers.FilesHandler).Methods("GET")
	router.HandleFunc("/upload", handlers.UploadHandler).Methods("POST")
}
