package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	IdUser       string  `json:"id_user"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	ListOfMangas []Manga `json:"mangas"`
}

var UserList []User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UserList)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get parameters
	//Loop through users and find with Id
	for _, item := range UserList {
		if item.IdUser == params["userid"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func GetUserId(users []User) string {
	if len(users) == 0 {
		return "1"
	}
	neoID, _ := strconv.Atoi(users[len(users)-1].IdUser)
	return strconv.Itoa(neoID + 1)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//Create user and set the correct Id
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.IdUser = GetUserId(UserList)

	user.ListOfMangas = MangaList
	MangaList = nil //reset mangas

	UserList = append(UserList, user)
	json.NewEncoder(w).Encode(user)
	fmt.Printf("user %v added\n", user.IdUser)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
