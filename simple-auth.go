package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jairoandre/simple-auth/model"
)

func postUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user model.User
	json.Unmarshal(reqBody, &user)
	user = addUser(user)
	json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := findAllUsers()
	json.NewEncoder(w).Encode(users)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func server() {
	fmt.Println("Starting at port 8080")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/user", postUser).Methods("POST")
	router.HandleFunc("/user", getUsers).Methods("GET")
	router.Handle("/redirect", http.RedirectHandler("https://github.com/jairoandre", 307))
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}

func main() {
	server()
}
