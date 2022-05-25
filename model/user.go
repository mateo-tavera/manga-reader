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
		if item.IdUser == params["uid"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	neoID, _ := strconv.Atoi(UserList[len(UserList)-1].IdUser)
	user.IdUser = strconv.Itoa(neoID + 1)

	UserList = append(UserList, user)
	json.NewEncoder(w).Encode(user)
	fmt.Printf("user %v added\n", neoID+1)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
