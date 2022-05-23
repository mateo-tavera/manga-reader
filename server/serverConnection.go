package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateo-tavera/manga/model"
)

func GetServerConnection() {
	//Init Router
	r := mux.NewRouter()

	//Route handlers / Endpoints
	r.HandleFunc("/api/mangas", model.GetMangas).Methods("GET")
	r.HandleFunc("/api/manga/{id}", model.GetManga).Methods("GET")
	r.HandleFunc("/api/manga", model.CreateManga).Methods("POST")
	r.HandleFunc("/api/manga/{id}", model.UpdateManga).Methods("PUT")
	r.HandleFunc("/api/manga/{id}", model.DeleteManga).Methods("DELETE")

	//Initilize the server
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
