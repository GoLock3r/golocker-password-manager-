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

var validated = true
var URI = "mongodb://localhost:27017"
var usernameglobal = ""

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
	usernameglobal = username
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
func home(w http.ResponseWriter, r *http.Request) {
	if validated {
		w.WriteHeader(http.StatusOK)
		var fileName = "home.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, usernameglobal)
		if err != nil {
			fmt.Println("Template execution error")

		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "well this isnt good your homepage should be here")
	}
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
		home(w, r)
		fmt.Println("should be a homepage")
	case "/home/display":
		readAll(w, r)
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
