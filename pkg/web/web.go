package web

import (
	"fmt"
	"golock3r/server/authtool"
	"golock3r/server/db"
	"golock3r/server/logger"
	"html/template"
	"net/http"
	"strings"
)

var validated = false
var URI = "mongodb://localhost:27017"
var Loggers *logger.Loggers

// func loadHTML(w http.ResponseWriter, r *http.Request, fileName string) {

// }

func login(w http.ResponseWriter, r *http.Request) {
	var fileName = "login.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Parse error")
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Template execution error")
	}
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	loggers := logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = loggers
	authtool.LoginFile = "logins.txt"
	db.Loggers = loggers

	username := r.FormValue("username")
	password := r.FormValue("password")

	validated = authtool.ValidateUser(username, string(password))
	if validated {
		var trimmedUser = strings.TrimSpace(username)
		db.Connect(trimmedUser)
		w.WriteHeader(http.StatusOK)
		var fileName = "login-submit.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, username)
		if err != nil {
			fmt.Println("Template execution error")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Login was unsuccessful, sit tight or try again who am I to tell you what to do.")
	}

}

func logout(w http.ResponseWriter, r *http.Request) {
	loggers := logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = loggers
	validated = db.CloseClientDB()
	if validated {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Your login was successful. Welcome to GoLock3r!")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Unexpected error. Logout was unsuccessful.")
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {

}

func message(w http.ResponseWriter, r *http.Request) {

}

func readAll(w http.ResponseWriter, r *http.Request) {

}

func search(w http.ResponseWriter, r *http.Request) {

}

func delete(w http.ResponseWriter, r *http.Request) {

}

func createEntry(w http.ResponseWriter, r *http.Request) {

}

func edit(w http.ResponseWriter, r *http.Request) {

}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		login(w, r)
		fmt.Println("A login page should be here")
	case "/login-submit":
		loginSubmit(w, r)
		fmt.Println("Submit login")
	case "/logout":
		logout(w, r)
		fmt.Println("Submit logout")
	case "/home":
		// var fileName = "home.html"
		// var username = "demo"
		// t, err := template.ParseFiles(fileName)
		// if err != nil {
		// 	fmt.Println("Parse error")
		// 	return
		// }
		// err = t.ExecuteTemplate(w, fileName, username)
		// if err != nil {
		// 	fmt.Println("Template execution error")
		// }

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
