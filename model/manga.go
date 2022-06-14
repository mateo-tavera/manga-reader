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
}

func GetMangaId(mangas []Manga) string {
	if len(mangas) == 0 {
		return "1"
	}
	neoID, _ := strconv.Atoi(mangas[len(mangas)-1].IdManga)
	return strconv.Itoa(neoID + 1)

}

func CreateManga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	currentUserId := params["userid"]
	var manga Manga
	_ = json.NewDecoder(r.Body).Decode(&manga)
	manga.IdManga = GetMangaId(MangaList)

	if currentUserId == params["userid"] {
		MangaList = append(MangaList, manga)
	} else {
		MangaList = nil
	}

	json.NewEncoder(w).Encode(manga)
	fmt.Printf("manga %v added\n", manga.IdManga)

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
