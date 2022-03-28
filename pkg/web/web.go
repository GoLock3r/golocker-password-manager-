package web

import (
	"fmt"
	"golock3r/server/authtool"
	"golock3r/server/db"
	"golock3r/server/logger"
	"html/template"
	"net/http"
)

var validated = false

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
	username := r.FormValue("username")
	password := r.FormValue("password")
	validated = authtool.ValidateUser(username, string(password))
	if validated {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Your login was successful. Welcome to GoLock3r! "+username)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Something went wrong")
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
		fmt.Fprintf(w, "Unexpected error logout was unsucessful")
	}
}

func message(w http.ResponseWriter, r *http.Request) {
	loggers := logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = loggers

}

func readAll(w http.ResponseWriter, r *http.Request) {

}

func search(w http.ResponseWriter, r *http.Request) {
	loggers := logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = loggers
}

func delete(w http.ResponseWriter, r *http.Request) {
	loggers := logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = loggers
}

func createUser(w http.ResponseWriter, r *http.Request) {
	loggers := logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = loggers
}

func createEntry(w http.ResponseWriter, r *http.Request) {
	loggers := logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = loggers
}

func edit(w http.ResponseWriter, r *http.Request) {
	loggers := logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = loggers
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
