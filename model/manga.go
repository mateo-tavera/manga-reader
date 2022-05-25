package model

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/mateo-tavera/manga/database"
)

type Manga struct {
	Title   string `json:"title"`
	Cover   string `json:"cover"`
	Website string `json:"website"`
	Chapter int    `json:"chapter"`
	IdManga string `json:"id_manga"`
}

var MangaList []Manga

func GetMangas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MangaList)
}

func GetManga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get parameters
	//Loop through mangas and find with Id
	for _, item := range MangaList {
		if item.IdManga == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Manga{})
}

func CreateManga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var manga Manga
	_ = json.NewDecoder(r.Body).Decode(&manga)
	neoID, _ := strconv.Atoi(MangaList[len(MangaList)-1].IdManga)
	manga.IdManga = strconv.Itoa(neoID + 1)

	MangaList = append(MangaList, manga)
	json.NewEncoder(w).Encode(manga)
	fmt.Printf("manga %v added\n", neoID+1)

	//Get db
	Db, err := database.GetDBConnection()
	if err != nil {
		log.Fatal("Cannot get DB connection", err)
	}

	//Save manga to db
	_, err = Db.Exec("INSERT INTO mangas (id, title, cover, website, chapter) VALUES (?, ?, ?, ?, ?)",
		manga.IdManga, manga.Title, manga.Cover, manga.Website, manga.Chapter)
	if err != nil {
		log.Fatal("Cannot execute query:", err)
	}
}

func DeleteManga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func UpdateManga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
