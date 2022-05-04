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
//Utilizes authtool package
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

//formats user inputed data to be put into the maps that are entered into the db
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
func readAll(w http.ResponseWriter, r *http.Request) bool {

	type Entry struct {
		Title       string
		Username    string
		Password    string
		PublicNote  string
		PrivateNote string
	}

	var entries []map[string]string
	var cards = ""

	if validated {
		// var display = ""
		var fileName = Path + "display.html"
		entries = db.ReadAll()
		for i := 0; i < len(entries); i++ {
			entries[i] = db.DecryptEntry(key, entries[i])
			var titleString = entries[i]["title"]
			var usernameString = entries[i]["username"]
			var passwordString = entries[i]["password"]
			var publicNoteString = entries[i]["public_note"]
			var privateNoteString = entries[i]["public_note"]
			cards += "<div class=\"col\"><div class=\"card shadow-sm\"><img src=\"...\" class=\"card-img-top\" alt=\"...\"><div class=\"card-body\"><h5 class=\"card-title\">" +
				"Title: " + titleString + "</h5><p class=\"card-text\">" +
				"Username: " + usernameString + "</p><p class=\"card-text\">" +
				"Password: " + passwordString + "</p><p class=\"card-text\">" +
				"Public Note: " + publicNoteString + "</p><p class=\"card-text\">" +
				"Private Note: " + privateNoteString + "</p></div></div></div>"
		}

		cards += "</div>"

		fmt.Println(cards)

		// display = formatEntryString(entries)
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}
		// Entry{entries[0]["title"], entries[0]["username"], entries[0]["password"], entries[0]["public_note"], entries[0]["private_note"]}
		err = t.ExecuteTemplate(w, "display.html", template.HTML(cards))
		if err != nil {
			fmt.Println("Template execution error")
			return false

		}
		w.WriteHeader(http.StatusOK)
		return true
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
		return false
	}
}

// Searches entry titles and usernames and displays the results
// for a validated user
func searchByTitle(w http.ResponseWriter, r *http.Request) bool {
	if validated {

		var fileName = Path + "searchTitle.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}
		err = t.ExecuteTemplate(w, "searchTitle.html", nil)
		if err != nil {
			fmt.Println("Template execution error")
			return false
		}
		return true
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
		return false
	}
}

// Gets form data and performs the search by title
func searchByTitle_submit(w http.ResponseWriter, r *http.Request) {
	if validated {

		var display = ""
		var fileName = Path + "searchByTitle.html"
		display = formatEntryString(db.ReadFromTitle(r.FormValue("title")))

		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, "searchByTitle.html", display)
		if err != nil {
			fmt.Println("Template execution error")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
	}
}

//serves page with html form that allows for search by username
func searchByUsername(w http.ResponseWriter, r *http.Request) bool {
	if validated {
		var fileName = Path + "searchUsername.html"
		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return false
		}
		err = t.ExecuteTemplate(w, "searchUsername.html", nil)
		if err != nil {
			fmt.Println("Template execution error")
			return false
		}
		return true
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Oh no maybe log in first")
		return false
	}
}

//grabs from data when form is submited and searches using user inputed username
func searchByUsername_submit(w http.ResponseWriter, r *http.Request) {

	if validated {
		var display = ""
		var fileName = Path + "searchByUsername.html"
		display = formatEntryString(db.ReadFromUsername(r.FormValue("username")))

		t, err := template.ParseFiles(fileName)
		if err != nil {
			fmt.Println("Parse error")
			return
		}
		err = t.ExecuteTemplate(w, "searchByUsername.html", display)
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
		fmt.Fprintf(w, "Oh no maybe log in first")
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
		fmt.Fprintf(w, "Oh no maybe log in first")
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
		fmt.Fprintf(w, "Oh no maybe log in first")
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
		fmt.Fprintf(w, "Oh no maybe log in first")
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
		fmt.Fprintf(w, "Oh no maybe log in first")
		return false
	}
}

//submits edit into active db
func edit_submit(w http.ResponseWriter, r *http.Request) bool {

	if validated {
		db.UpdateEntry(r.FormValue("title"), r.FormValue("update_key"), r.FormValue("update_value"))
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
		fmt.Fprintf(w, "Oh no maybe log in first")
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
		fmt.Fprintf(w, "Oh no maybe log in first")
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
		w.WriteHeader(http.StatusNotFound)
	}
}

// Creates an instance of the web server. Listens on port 8010
func Run() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8010", nil)
}
