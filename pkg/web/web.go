package web

import (
	"fmt"
	"net/http"
)

var validated = false

func login(w http.ResponseWriter, r *http.Request) {

}

func loginSubmit(w http.ResponseWriter, r *http.Request) {

}

func logout(w http.ResponseWriter, r *http.Request) {

}

func message(w http.ResponseWriter, r *http.Request) {

}

func readAll(w http.ResponseWriter, r *http.Request) {

}

func search(w http.ResponseWriter, r *http.Request) {

}

func delete(w http.ResponseWriter, r *http.Request) {

}

func createUser(w http.ResponseWriter, r *http.Request) {

}

func createEntry(w http.ResponseWriter, r *http.Request) {

}

func edit(w http.ResponseWriter, r *http.Request) {

}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Println("A login page should be here")
	case "/login-submit":
		fmt.Println("Submit login")
	case "/logout":
		fmt.Println("Submit logout")
	case "/home":
		fmt.Println("Users homepage")
	case "/home/display":
		fmt.Println("Display all db entries")
	case "/home/search":
		fmt.Println("Search here")
	case "/home/delete":
		fmt.Println("Delete here")
	case "/home/create":
		fmt.Println("Create here")
	case "/home/edit":
		fmt.Println("Edit here")
	default:
		fmt.Println("Path not found?")
	}
}

func Run() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8010", nil)
}
