package main

import (
	"fmt"
	"log"

	"github.com/mateo-tavera/manga/database"
	"github.com/mateo-tavera/manga/model"
	"github.com/mateo-tavera/manga/server"
)

func main() {

	//Set db connection
	Db, err := database.GetDBConnection()
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	fmt.Println("connected to database")
	defer Db.Close()

	//Delete previus data
	_, err = Db.Exec("TRUNCATE mangas")
	if err != nil {
		log.Fatal("Cannot execute query:", err)
	}

	prueba := []string{}

	fmt.Println(len(prueba))
	//Data to create a manga manually

	model.MangaList = append(model.MangaList, model.Manga{
		Title:   "Boku No Hero",
		Cover:   "cover",
		Website: "website",
		Chapter: 5,
		IdManga: "1"})

	model.MangaList = append(model.MangaList, model.Manga{
		Title:   "One Punch Man",
		Cover:   "cover",
		Website: "website",
		Chapter: 1,
		IdManga: "2"})

	//Data to create a user manually

	model.UserList = append(model.UserList, model.User{
		IdUser:       "1",
		FirstName:    "Pablo",
		LastName:     "Bernal",
		ListOfMangas: nil})

	database.GetDBConnection()
	server.GetServerConnection()

}
