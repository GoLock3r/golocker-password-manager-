package web

import (
	"fmt"
	"golock3r/server/authtool"
	"golock3r/server/crypt"
	"golock3r/server/db"
	"golock3r/server/logger"
	"html/template"
	"net/http"
	"strings"
)

var Loggers *logger.Loggers

var valid_username = ""

var validated = false

var key []byte

var Path = "web/assets/"

var Url = "http://localhost:8010"

// Serve a login page to the user and pass credentials off to the
// loginSubmit function to verify these credentials
func login(w http.ResponseWriter, r *http.Request) bool {
	if !validated {
		var fileName = Path + "login.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}
		err = t.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			fmt.Println("Template execution error")
			return false
		}
		return true
	} else {
		var fileName = Path + "redirect.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}
		err = t.ExecuteTemplate(w, "redirect.html", Url+"/home")
		if err != nil {
			fmt.Println("Template execution error")
			return false
		}
	}
	return true
}

// Process user credentials given from the login function.
// Utilizes authtool package functionality to validate credentials.
func loginSubmit(w http.ResponseWriter, r *http.Request) bool {
	username := r.FormValue("username")
	password := r.FormValue("password")
	key = authtool.GetKey(username, password)
	validated = authtool.ValidateUser(username, string(password))
	if validated {
		var trimmedUser = strings.TrimSpace(username)
		db.Loggers = Loggers
		db.Connect(trimmedUser)
		w.WriteHeader(http.StatusOK)
		var fileName = Path + "redirect.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}
		err = t.ExecuteTemplate(w, "redirect.html", Url+"/home")
		if err != nil {
			fmt.Println(err, "Template execution error")
			return false
		}
		valid_username = username
		return true
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Login was unsuccessful, sit tight or try again; who am I to tell you what to do?")
		return false
	}

}

// Strip validation from the users database and dissconects from said data base\
func logout(w http.ResponseWriter, r *http.Request) bool {

	if validated {
		db.CloseClientDB()
		validated = false
		key = nil
		valid_username = ""

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Your logout was successful. From us at GoLock3r, goodbye!")
		return true
	} else {
		w.WriteHeader(http.StatusNotFound)
		Loggers.LogError.Println("Unexpected error. Logout was unsuccessful.")
		fmt.Fprintf(w, "Unexpected error. Logout was unsuccessful.")
		return false
	}
}

// Creates a new valid user account
// Utilizes authtool package
func createUser(w http.ResponseWriter, r *http.Request) bool {
	username := r.FormValue("username")
	password := r.FormValue("password")
	var userCreated = authtool.CreateUser(username, password)
	if userCreated {
		w.WriteHeader(http.StatusOK)
		var fileName = Path + "redirect.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}
		err = t.ExecuteTemplate(w, "redirect.html", Url+"/")
		if err != nil {
			fmt.Println(err, "Template execution error")
			return false
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Unable to create account ")
		return false
	}
	return true
}

// Parses entry data from various read methods into html cards
func parseCards(entryTable []map[string]string, callingMethod string) string {
	var cards = ""
	if entryTable == nil {
		if callingMethod == "readAll" {
			cards += "<h>No entries found. </br>Try creating a new entry!</h>"
		} else {
			cards += "<h>Your search didn't return anything. </br>Why don't you try again?</h>"
		}
	} else {
		for _, entry := range entryTable {
			entry = db.DecryptEntry(key, entry)
			cards += "<div class=\"col\"><div class=\"card shadow-sm\"><img src=\"...\" class=\"card-img-top\" alt=\"...\"><div class=\"card-body\"><h5 class=\"card-title\">" +
				"Title: " + entry["title"] + "</h5><p class=\"card-text\">" +
				"Username: " + entry["username"] + "</p><p class=\"card-text\">" +
				"Password: " + entry["password"] + "</p><p class=\"card-text\">" +
				"Private Note: " + entry["private_note"] + "</p><p class=\"card-text\">" +
				"Public Note: " + entry["public_note"] + "</p></div></div></div>"
		}
	}
	return cards
}

// Reads all database entries for a validated user and
// displays all of the entries in user database
func readAll(w http.ResponseWriter, r *http.Request) bool {

	if validated {

		var fileName = Path + "display.html"
		var cards = parseCards(db.ReadAll(), "readAll")

		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}

		err = t.ExecuteTemplate(w, "display.html", template.HTML(cards))
		if err != nil {
			fmt.Println("Template execution error")
			return false
		}
		// w.WriteHeader(http.StatusOK) -- Note: Output on command line says this call is superfluous; removing it broke nothing
		return true
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "You are not logged in.")
		return false
	}
}

// Gets form data and performs the search by title
func search(w http.ResponseWriter, r *http.Request) bool{

	var cards = ""

	if validated {

		var fileName = Path + "display.html"
		if r.FormValue("searchType") == "title" {
			cards = parseCards(db.ReadFromTitle(r.FormValue("searchString")), "search")
		} else {
			cards = parseCards(db.ReadFromUsername(r.FormValue("searchString")), "search")
		}

		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}
		err = t.ExecuteTemplate(w, "display.html", template.HTML(cards))
		if err != nil {
			fmt.Println("Template execution error")
			return false
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "You are not logged in.")
		return false
	}
	return true
}

// Deletes an entry from the database for a validated user
func delete_submit(w http.ResponseWriter, r *http.Request) {

	if validated {
		var fileName = Path + "redirect.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, "redirect.html", Url+"/home")
		if err != nil {
			fmt.Println("Template execution error")
		}
		title := r.FormValue("title")
		db.DeleteEntry(title)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "You are not logged in.")
	}
}

//Servers form for user to enter the entry that they would like to delete
func delete(w http.ResponseWriter, r *http.Request) bool {

	if validated {
		var fileName = Path + "delete.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}

		err = t.ExecuteTemplate(w, "delete.html", nil)
		if err != nil {
			fmt.Println("Template execution error")
			return false
		}
		return true
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "You are not logged in.")
		return false
	}
}

// Creates a new entry to be securely stored on the database for
// a validated user
func createEntry(w http.ResponseWriter, r *http.Request) bool {

	if validated {
		var fileName = Path + "create.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}
		err = t.ExecuteTemplate(w, "create.html", nil)
		if err != nil {
			fmt.Println("Template execution error")
			return false
		}
		return true

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "You are not logged in.")
		return false
	}
}

//submits entry into db with encryption
func createEntrySubmit(w http.ResponseWriter, r *http.Request) bool {

	if validated {
		entry := map[string]string{
			"title":        r.FormValue("title"),
			"password":     r.FormValue("password"),
			"username":     r.FormValue("username"),
			"private_note": r.FormValue("private_note"),
			"public_note":  r.FormValue("public_note"),
		}
		entry = db.EncryptEntry(key, entry)
		db.WriteEntry(entry)
		var fileName = Path + "redirect.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}
		err = t.ExecuteTemplate(w, "redirect.html", Url+"/home")
		if err != nil {
			fmt.Println("Template execution error")
			return false
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "You are not logged in.")
		return false
	}
	return true
}

// Edits an entry for a validated user
func edit(w http.ResponseWriter, r *http.Request) bool {

	if validated {
		var fileName = Path + "edit.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}
		err = t.ExecuteTemplate(w, "edit.html", nil)
		if err != nil {
			fmt.Println("Template execution error")
			return false
		}
		return true
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "You are not logged in.")
		return false
	}
}

//submits edit into active db
func edit_submit(w http.ResponseWriter, r *http.Request) bool {

	var updateValue = ""

	if validated {
		if r.FormValue("updateField") == "password" || r.FormValue("updateField") == "private_note" {
			updateValue = crypt.EncryptStringToHex(key, r.FormValue("updateValue"))
		} else {
			updateValue = r.FormValue("updateValue")
		}

		db.UpdateEntry(r.FormValue("title"), r.FormValue("updateField"), updateValue)
		var fileName = Path + "redirect.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}
		err = t.ExecuteTemplate(w, "redirect.html", Url+"/home")
		if err != nil {
			fmt.Println("Template execution error")
			return false
		}

	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "You are not logged in.")
		return false
	}
	return true
}

// Display the homepage for a validated user
func home(w http.ResponseWriter, r *http.Request) bool {

	if validated {
		w.WriteHeader(http.StatusOK)
		var fileName = Path + "home.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}
		err = t.ExecuteTemplate(w, "home.html", valid_username)
		if err != nil {
			fmt.Println("Template execution error")
			return false
		}
		return true
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "You are not logged in.")
		return false
	}

}

// Handle the navigation logic for the server's resources and functions
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		login(w, r)
	case "":
		login(w, r)
	case "/login-submit":
		loginSubmit(w, r)
	case "/logout":
		logout(w, r)
	case "/home":
		home(w, r)
	case "/home/display":
		readAll(w, r)
	case "/home/search":
		search(w, r)
	case "/home/delete":
		delete(w, r)
	case "/home/delete-submit":
		delete_submit(w, r)
	case "/home/create":
		createEntry(w, r)
	case "/home/create-Submit":
		createEntrySubmit(w, r)
	case "/home/edit":
		edit(w, r)
	case "/home/edit-submit":
		edit_submit(w, r)
	case "/createUser":
		createUser(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

// Creates an instance of the web server. Listens on port 8010
func Run() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8010", nil)
}
