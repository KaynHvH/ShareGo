package main

import (
	"ShareGo/api/config"
	"ShareGo/api/db"
	"ShareGo/api/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	db.InitDB()

	router := mux.NewRouter()
	routes.InitRoutes(router)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	log.Println("Listening on " + cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, router)
	if err != nil {
		log.Fatal(err)
		return
	}
}
