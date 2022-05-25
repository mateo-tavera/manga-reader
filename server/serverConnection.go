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
	//User routes
	r.HandleFunc("/api/users/", model.GetUsers).Methods("GET")
	r.HandleFunc("/api/user/{uid}", model.GetUser).Methods("GET")
	r.HandleFunc("/api/users/", model.CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{uid}", model.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{uid}", model.DeleteUser).Methods("DELETE")
	//Manga routes
	r.HandleFunc("/api/users/{uid}/mangas", model.GetMangas).Methods("GET")
	r.HandleFunc("/api/users/{uid}/manga/{id}", model.GetManga).Methods("GET")
	r.HandleFunc("/api/users/{uid}/manga", model.CreateManga).Methods("POST")
	r.HandleFunc("/api/users/{uid}/manga/{id}", model.UpdateManga).Methods("PUT")
	r.HandleFunc("/api/users/{uid}/manga/{id}", model.DeleteManga).Methods("DELETE")

	//Initilize the server
	fmt.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
