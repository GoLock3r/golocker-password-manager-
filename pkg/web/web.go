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
var key []byte

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
	key = authtool.GetKey(username, password)
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
		fmt.Fprintf(w, "Login was unsuccessful, sit tight or try again; who am I to tell you what to do?")
	}

}

// Strip validation from the users database and dissconects from said data base\
func logout(w http.ResponseWriter, r *http.Request) {
	loggers := logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = loggers

	if validated {
		db.CloseClientDB()
		validated = false
		key = nil
		valid_username = ""

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Your logout was successful. From us at GoLock3r, goodbye!")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Unexpected error. Logout was unsuccessful.")
	}
}

// Creates a new valid user account
//Utilizes authtool package
func createUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	var userCreated = authtool.CreateUser(username, password)
	if userCreated {
		w.WriteHeader(http.StatusOK)
		var fileName = "createUser-submit.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, username)
		if err != nil {
			fmt.Println("Template execution error")
		}
		fmt.Fprintf(w, "Account was created successfully")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Unable to create acount ")
	}
}

func formatEntryString(entryTable []map[string]string) string {
	var display = ""
	if entryTable == nil {
		display = "No entries"
	}
	for _, entry := range entryTable {
		display += "Website: " + entry["title"] + "\n"
		display += "Username: " + entry["username"] + "\n"
		display += "Password: " + entry["password"] + "\n"
		display += "Public note: " + entry["public_note"] + "\n"
		display += "Private note: " + entry["private_note"] + "\n"
		display += "\n"
	}
	return display
}

// Reads all database entries for a validated user
// displays all of the entries in user database
func readAll(w http.ResponseWriter, r *http.Request) {
	var entries []map[string]string
	if validated {
		var display = ""
		var fileName = "display.html"
		entries = db.ReadAll()
		for i := 0; i < len(entries); i++ {
			entries[i] = db.DecryptEntry(key, entries[i])
		}
		display = formatEntryString(entries)
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, display)
		if err != nil {
			fmt.Println("Template execution error")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
	}
}

// Searches entry titles and usernames and displays the results
// for a validated user
func searchByTitle(w http.ResponseWriter, r *http.Request) {
	if validated {

		var fileName = "searchTitle.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, nil)
		if err != nil {
			fmt.Println("Template execution error")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
	}
}
func searchByTitle_submit(w http.ResponseWriter, r *http.Request) {
	if validated {

		var display = ""
		var fileName = "searchByTitle.html"
		display = formatEntryString(db.ReadFromTitle(r.FormValue("title")))

		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, display)
		if err != nil {
			fmt.Println("Template execution error")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
	}
}

func searchByUsername(w http.ResponseWriter, r *http.Request) {
	if validated {
		var fileName = "searchUsername.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, nil)
		if err != nil {
			fmt.Println("Template execution error")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
	}
}
func searchByUsername_submit(w http.ResponseWriter, r *http.Request) {
	if validated {
		var display = ""
		var fileName = "searchByUsername.html"
		display = formatEntryString(db.ReadFromUsername(r.FormValue("username")))

		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, display)
		if err != nil {
			fmt.Println("Template execution error")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
	}
}

// Deletes an entry from the database for a validated user
func delete_submit(w http.ResponseWriter, r *http.Request) {
	if validated {
		var fileName = "delete-submit.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, nil)
		if err != nil {
			fmt.Println("Template execution error")
		}
		title := r.FormValue("title")
		db.DeleteEntry(title)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	if validated {
		var fileName = "delete.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, nil)
		if err != nil {
			fmt.Println("Template execution error")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
	}
}

// Creates a new entry to be securely stored on the database for
// a validated user
func createEntry(w http.ResponseWriter, r *http.Request) {
	if validated {
		var fileName = "create.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, nil)
		if err != nil {
			fmt.Println("Template execution error")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
	}
}

func createEntrySubmit(w http.ResponseWriter, r *http.Request) {
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
		var fileName = "createsubmit.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, nil)
		if err != nil {
			fmt.Println("Template execution error")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
	}
}

// Edits an entry for a validated user
func edit(w http.ResponseWriter, r *http.Request) {
	if validated {
		var fileName = "edit.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, fileName, nil)
		if err != nil {
			fmt.Println("Template execution error")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
	}
}

func edit_submit(w http.ResponseWriter, r *http.Request) {
	if validated {
		db.UpdateEntry(r.FormValue("title"), r.FormValue("update_key"), r.FormValue("update_value"))
		if validated {
			var fileName = "edit-submit.html"
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
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
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
		fmt.Fprintf(w, "Oh no maybe log in first")
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
	case "/home/searchTitle":
		searchByTitle(w, r)
	case "/home/searchTitle-Submit":
		searchByTitle_submit(w, r)
	case "/home/searchUser":
		searchByUsername(w, r)
	case "/home/searchUser-Submit":
		searchByUsername_submit(w, r)
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
		//fmt.Println("Path not found?")
	}
}

// Creates an instance of the web server. Listens on port 8010
func Run() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8010", nil)
}
