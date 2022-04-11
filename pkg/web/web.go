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

var Loggers *logger.Loggers

var valid_username = ""
var validated = false

// Serve a login page to the user and pass credentials off to the
// loginSubmit function to verify these credentials
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

// Process user credentials given from the login function.
// Utilizes authtool package functionality to validate credentials.
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
		valid_username = username

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Login was unsuccessful, sit tight or try again who am I to tell you what to do.")
	}

}

// Strip validation from the users current session redirect elsewhere
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

// Creates a new valid user account
func createUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	var userCreated = authtool.CreateUser(username, password)

	if userCreated == true{
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "New user created return to login page ")
	}else{
		w.WriteHeader(http.StatusNotFound)
	    fmt.Fprint(w, "Unable to create new user try again")
	}
}

// Reads all database entries for a validated user
func readAll(w http.ResponseWriter, r *http.Request) {
if validated{
	rA = db.ReadAll()
}
}

// Searches entry titles and usernames and displays the results
// for a validated user
func searchByTitle(w http.ResponseWriter, r *http.Request) {
if validated{
	title := r.FormValue("title")
	var sBT = db.ReadFromTitle(title)
}
}
func searchByUsername(w http.ResponseWriter, r *http.Request) {
	if validated{
	username := r.FormValue("username")
	var sBU = db.ReadFromUsername(username)
	}
}

// Deletes an entry from the database for a validated user
func delete(w http.ResponseWriter, r *http.Request) {
if validated{
	title := r.FormValue("title")
	db.DeleteEntry(title)
}
}

// Creates a new entry to be securely stored on the database for
// a validated user
func createEntry(w http.ResponseWriter, r *http.Request) {
if validated{
	db.WriteEntry()
}
}

// Edits an entry for a validated user
func edit(w http.ResponseWriter, r *http.Request) {
if validated{
	db.UpdateEntry()

}
}

// Display the homepage for a validated user
func home(w http.ResponseWriter, r *http.Request) {
	if validated {
		w.WriteHeader(http.StatusOK)
		var fileName = "home.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, valid_username)
		if err != nil {
			fmt.Println("Template execution error")

		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "well this isn't good your homepage should be here")
	}
}

// Handle the navigation logic for the server's resources and functions
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
	case "/createUser":
		createUser(w,r)
		fmt.Print("just be a create user page")
	default:
		fmt.Println("Path not found?")
	}
}

// Creates an instance of the web server. Listens on port 8010
func Run() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8010", nil)
}
