package web

import (
	"golock3r/server/authtool"
	"golock3r/server/db"
	"golock3r/server/logger"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func removeFiles() {
	os.Remove("logins.txt")
	os.Remove("testlogs.txt")
}
func TestLandingPage(t *testing.T) {
	Path = "assets/"
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	var landingpage = login(w, req)
	if !landingpage {
		t.Error("expected to have the landing page shown landing page didnt show ")
	}
}
func TestLandingPageWrongPath(t *testing.T) {
	Path = "something/"
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	var landingpage = login(w, req)
	if landingpage {
		t.Error("expected to have the not landing page shown landing page didnt show ")
	}
}
func TestLandingPagevalidated(t *testing.T) {
	Path = "assets/"
	validated = true
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	var landingpage = login(w, req)
	if !landingpage {
		t.Error("expected to have the landing page shown landing page didnt show ")
	}
}

func TestLoginSubmit(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	authtool.CreateUser("test_username", "test_password")
	Path = "assets/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/login-submit?username=test_username&password=test_password", nil)
	loginSubmit := loginSubmit(w, req)
	if !loginSubmit {
		t.Error("login should have submitted succesfully it didnt")
	}

}
func TestLoginSubmitinvalidLogIn(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	authtool.CreateUser("test_username", "test_password")
	Path = "assets/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/login-submit?username=test_username1&password=test_password", nil)
	loginSubmit := loginSubmit(w, req)
	if loginSubmit {
		t.Error("login should have not submitted succesfully it didnt")
	}

}
func TestLoginSubmitWrongPath(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	authtool.CreateUser("test_username", "test_password")
	Path = "something/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/login-submit?username=test_username&password=test_password", nil)
	loginSubmit := loginSubmit(w, req)
	if loginSubmit {
		t.Error("login should have submitted succesfully it didnt")
	}

}

// invalid login test
func TestCreateUser(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	Path = "assets/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/createUser?username=test_username4&password=test_password1", nil)
	loginSubmit := createUser(w, req)
	if !loginSubmit {
		t.Error("login should have submitted succesfully it didnt")
	}

}
func TestCreateUserexisting(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	Path = "assets/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/createUser?username=test_username4&password=test_password1", nil)
	loginSubmit := createUser(w, req)
	if loginSubmit {
		t.Error("shouldnt have created user")
	}

}
func TestCreateUserWrongPath(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	Path = "something/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/createUser?username=test_username4&password=test_password1", nil)
	loginSubmit := createUser(w, req)
	if loginSubmit {
		t.Error("login should have submitted succesfully it didnt")
	}

}
// create user with existing user
func TestLogout(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	//authtool.Loggers = Loggers
	db.Loggers = Loggers
	db.Connect("demo")
	validated = true
	req := httptest.NewRequest(http.MethodGet, "/logout", nil)
	w := httptest.NewRecorder()
	var logout = logout(w, req)
	if !logout {
		t.Error("logout unsuccessful", w)
	}

}

// testLogout when validated == false
func TestLogoutInvalidated(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	//authtool.Loggers = Loggers
	db.Loggers = Loggers
	db.Connect("demo")
	validated = false
	req := httptest.NewRequest(http.MethodGet, "/logout", nil)
	w := httptest.NewRecorder()
	var logout = logout(w, req)
	if logout {
		t.Error("logout successful")
	}

}

func TestReadall(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	validated = true
	Path = "assets/"
	db.Connect("test")
	req := httptest.NewRequest(http.MethodGet, "/home/display", nil)
	w := httptest.NewRecorder()
	var readAll = readAll(w, req)
	if !readAll {
		t.Error("unable to read all")

	}
}
func TestReadallWrongPath(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	validated = true
	Path = "something/"
	db.Connect("test")
	req := httptest.NewRequest(http.MethodGet, "/home/display", nil)
	w := httptest.NewRecorder()
	var readAll = readAll(w, req)
	if readAll {
		t.Error("unable to read all")

	}
}

func TestReadallInvalidated(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	validated = false
	Path = "assets/"
	db.Connect("test")
	req := httptest.NewRequest(http.MethodGet, "/home/display", nil)
	w := httptest.NewRecorder()
	var readAll = readAll(w, req)
	if readAll {
		t.Error("shouldnt be able to read all")

	}
}
func TestSearchByTitle(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	Path = "assets/"
	validated = true
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/searchTitle", nil)
	w := httptest.NewRecorder()
	var landingpage = searchByTitle(w, req)
	if !landingpage {
		t.Error("expected to have the search by title page shown")
	}

}
func TestSearchByTitleWrongPath(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	Path = "something/"
	validated = true
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/searchTitle", nil)
	w := httptest.NewRecorder()
	var landingpage = searchByTitle(w, req)
	if landingpage {
		t.Error("expected to have the search by title page shown")
	}

}
func TestSearchByTitleInvalidated(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	Path = "assets/"
	validated = false
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/searchTitle", nil)
	w := httptest.NewRecorder()
	var landingpage = searchByTitle(w, req)
	if landingpage {
		t.Error("title page should be shown")
	}

}
func TestSearchByTitleSubmit(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = true
	Path = "assets/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/login-submit?searchType=title&username=test_username&password=test_password", nil)
	loginSubmit(w, req)
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/home/searchTitle-Submit?title=test", nil)
	loginSubmit := searchByTitle(w, req)
	if !loginSubmit {
		t.Error("edit should have submitted succesfully it didnt")
	}
}
func TestSearchByTitleSubmitWrongPath(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = true
	Path = "something/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/login-submit?searchType=title&username=test_username&password=test_password", nil)
	loginSubmit(w, req)
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/home/searchTitle-Submit?title=test", nil)
	loginSubmit := searchByTitle(w, req)
	if loginSubmit {
		t.Error("edit should have submitted succesfully it didnt")
	}
}

func TestSearchByTitleSubmitInvalidated(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	Path = "assets/"
	validated = false
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/home/searchTitle-Submit?title=test", nil)
	loginSubmit := searchByTitle(w, req)
	if loginSubmit {
		t.Error("edit should not have submitted succesfully it didnt")
	}
}
func TestSearchByUsername(t *testing.T) {
	Path = "assets/"
	validated = true
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/searchUser", nil)
	w := httptest.NewRecorder()
	var landingpage = searchByUsername(w, req)
	if !landingpage {
		t.Error("expected to have the search by username page shown")
	}
}
func TestSearchByUsernameWrongPath(t *testing.T) {
	Path = "something/"
	validated = true
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/searchUser", nil)
	w := httptest.NewRecorder()
	var landingpage = searchByUsername(w, req)
	if landingpage {
		t.Error("expected to have the search by username page shown")
	}
}
func TestSearchByUsernameInvalidated(t *testing.T) {
	Path = "assets/"
	validated = false
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/searchUser", nil)
	w := httptest.NewRecorder()
	var landingpage = searchByUsername(w, req)
	if landingpage {
		t.Error("expected to not have the search by username page shown")
	}
}
func TestSearchByUsernameSubmit(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = true
	Path = "assets/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/login-submit?searchType=username&username=test_username&password=test_password", nil)
	loginSubmit(w, req)
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/home/searchUser-Submit?username=test", nil)
	loginSubmit := searchByUsername(w, req)
	if !loginSubmit {
		t.Error("edit should have submitted succesfully it didnt")
	}
}
func TestSearchByUsernameSubmitWrongPath(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = true
	Path = "something/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/login-submit?searchType=username&username=test_username&password=test_password", nil)
	loginSubmit(w, req)
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/home/searchUser-Submit?username=test", nil)
	loginSubmit := searchByUsername(w, req)
	if loginSubmit {
		t.Error("edit should have submitted succesfully it didnt")
	}
}

func TestSearchByUsernameSubmitInvalidated(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = false
	Path = "assets/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/home/searchUser-Submit?username=test", nil)
	loginSubmit := searchByUsername(w, req)
	if loginSubmit {
		t.Error("edit should not have submitted succesfully it didnt")
	}
}

func TestDelete(t *testing.T) {
	Path = "assets/"
	validated = true
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/delete", nil)
	w := httptest.NewRecorder()
	var landingpage = delete(w, req)
	if !landingpage {
		t.Error("expected to have the delete page shown")
	}
}
func TestDeleteWrongPath(t *testing.T) {
	Path = "something/"
	validated = true
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/delete", nil)
	w := httptest.NewRecorder()
	var landingpage = delete(w, req)
	if landingpage {
		t.Error("expected to have the delete page shown")
	}
}
func TestDeleteInvalidated(t *testing.T) {
	Path = "assets/"
	Loggers = logger.CreateLoggers("testlogs.txt")
	validated = false
	req := httptest.NewRequest(http.MethodGet, "/home/delete", nil)
	w := httptest.NewRecorder()
	var landingpage = delete(w, req)
	if landingpage {
		t.Error("expected to not have the delete page shown")
	}
}

func TestDeleteSubmit(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = true
	Path = "assets/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/login-submit?username=test_username&password=test_password", nil)
	loginSubmit(w, req)
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/home/delete-submit?title=test&update_key=title&update_value=test2", nil)
	loginSubmit := delete(w, req)
	if !loginSubmit {
		t.Error("edit should have submitted succesfully it didnt")
	}
}
func TestDeleteSubmitWrongPath(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = true
	Path = "something/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/login-submit?username=test_username&password=test_password", nil)
	loginSubmit(w, req)
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/home/delete-submit?title=test&update_key=title&update_value=test2", nil)
	loginSubmit := delete(w, req)
	if loginSubmit {
		t.Error("edit should have submitted succesfully it didnt")
	}
}
func TestDeleteSubmitInvalidated(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = false
	Path = "assets/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/home/delete-submit?title=test&update_key=title&update_value=test2", nil)
	loginSubmit := delete(w, req)
	if loginSubmit {
		t.Error("edit should not have submitted succesfully it didnt")
	}
}
func TestCreateEntry(t *testing.T) {
	Path = "assets/"
	validated = true
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/create", nil)
	w := httptest.NewRecorder()
	var landingpage = createEntry(w, req)
	if !landingpage {
		t.Error("expected to have the create entry page shown")
	}
}
func TestCreateEntryWrongPath(t *testing.T) {
	Path = "something/"
	validated = true
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/create", nil)
	w := httptest.NewRecorder()
	var landingpage = createEntry(w, req)
	if landingpage {
		t.Error("expected to have the create entry page shown")
	}
}
func TestCreateEntryInvalidated(t *testing.T) {
	Path = "assets/"
	validated = false
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/create", nil)
	w := httptest.NewRecorder()
	var landingpage = createEntry(w, req)
	if landingpage {
		t.Error("expected to have the create entry page shown")
	}
}

func TestCreateEntrySubmit(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = true
	Path = "assets/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/login-submit?username=test_username&password=test_password", nil)
	loginSubmit(w, req)
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/home/create-Submit?title=test&password=test&username=test&private_note=test&public_note=test", nil)
	loginSubmit := createEntrySubmit(w, req)
	if !loginSubmit {
		t.Error("create should have submitted succesfully it didnt")
	}
}
func TestCreateEntrySubmitWrongPath(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = true
	Path = "something/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/login-submit?username=test_username&password=test_password", nil)
	loginSubmit(w, req)
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/home/create-Submit?title=test&password=test&username=test&private_note=test&public_note=test", nil)
	loginSubmit := createEntrySubmit(w, req)
	if loginSubmit {
		t.Error("create should have submitted succesfully it didnt")
	}
}
func TestCreateEntrySubmitInvalidated(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = false
	Path = "assets/"

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/home/create-Submit?title=test&password=test&username=test&private_note=test&public_note=test", nil)
	loginSubmit := createEntrySubmit(w, req)
	if loginSubmit {
		t.Error("create should have submitted succesfully it didnt")
	}
}
func TestEdit(t *testing.T) {
	Path = "assets/"
	validated = true
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/edit", nil)
	w := httptest.NewRecorder()
	var landingpage = edit(w, req)
	if !landingpage {
		t.Error("expected to have the edit page didnt show ")
	}
}
func TestEditWrongPath(t *testing.T) {
	Path = "something/"
	validated = true
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/edit", nil)
	w := httptest.NewRecorder()
	var landingpage = edit(w, req)
	if landingpage {
		t.Error("expected to have the edit page didnt show ")
	}
}
func TestEditInvalidated(t *testing.T) {
	Path = "assets/"
	validated = false
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/edit", nil)
	w := httptest.NewRecorder()
	var landingpage = edit(w, req)
	if landingpage {
		t.Error("expected to have the edit page didnt show ")
	}
}

func TestEditSubmit(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = true
	Path = "assets/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/login-submit?username=test_username&password=test_password", nil)
	loginSubmit(w, req)
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/home/edit-submit?title=test&update_key=title&update_value=test2", nil)
	loginSubmit := edit_submit(w, req)
	if !loginSubmit {
		t.Error("edit should have submitted succesfully it didnt")
	}

}
func TestEditSubmitWrongPath(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = true
	Path = "something/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/login-submit?username=test_username&password=test_password", nil)
	loginSubmit(w, req)
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodGet, "/home/edit-submit?title=test&update_key=title&update_value=test2", nil)
	loginSubmit := edit_submit(w, req)
	if loginSubmit {
		t.Error("edit should have submitted succesfully it didnt")
	}

}
func TestEditSubmitInvalidated(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Loggers = Loggers
	authtool.Loggers = Loggers
	validated = false
	Path = "assets/"
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/home/edit-submit?title=test&update_key=title&update_value=test2", nil)
	loginSubmit := edit_submit(w, req)
	if loginSubmit {
		t.Error("edit should have submitted succesfully it didnt")
	}

}

func TestParseCardsReadAll(t *testing.T){
	entry := map[string]string{
		"title":        "title",
		"password":     "password",
		"username":     "username",
		"private_note": "private_note",
		"public_note":  "public_note",
	}
	var results_map []map[string]string

	results_map = append(results_map, entry)
var ans =parseCards(results_map, "readAll")

	if ans == ""{
		t.Error("expected a response")
	}
}


func TestParseCardsreadAllempty(t *testing.T){
	var results_map []map[string]string
	var ans =parseCards(results_map, "readAll")
	if ans != "<h>No entries found. </br>Try creating a new entry!</h>"{
		t.Error("expected a response of <h>No entries found. </br>Try creating a new entry!</h> go t", ans)
	}
}
func TestParseCardsElse(t *testing.T){
	entry := map[string]string{
		"title":        "title",
		"password":     "password",
		"username":     "username",
		"private_note": "private_note",
		"public_note":  "public_note",
	}
	var results_map []map[string]string

	results_map = append(results_map, entry)
var ans =parseCards(results_map, "anything")

	if ans == ""{
		t.Error("expected a response")
	}
}
func TestParseCardselseempty(t *testing.T){
	var results_map []map[string]string
	var ans =parseCards(results_map, "else")
	if ans != "<h>Your search didn't return anything. </br>Why don't you try again?</h>" {
		t.Error("expected a response of <h>Your search didn't return anything. </br>Why don't you try again?</h> got ", ans)
	}
}
func TestHome(t *testing.T) {
	Path = "assets/"
	validated = true
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home", nil)
	w := httptest.NewRecorder()
	var landingpage = home(w, req)
	if !landingpage {
		t.Error("expected to have the home page shown landing page didnt show ")
	}
	removeFiles()
}
func TestHomeWrongPath(t *testing.T) {
	Path = "something/"
	validated = true
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home", nil)
	w := httptest.NewRecorder()
	var landingpage = home(w, req)
	if landingpage {
		t.Error("expected to have the home page shown landing page didnt show ")
	}
	removeFiles()
}
func TestHomeInvalidated(t *testing.T) {
	Path = "assets/"
	validated = false
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home", nil)
	w := httptest.NewRecorder()
	var landingpage = home(w, req)
	if landingpage {
		t.Error("expected to have the home page shown landing page didnt show ")
	}
	removeFiles()
}
func TestHandlercase1(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}

func TestHandlercase2(t *testing.T) { // needs db access
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	authtool.CreateUser("test_username", "test_password")
	req := httptest.NewRequest(http.MethodGet, "/login-submit?username=test_username&password=test_password", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}
func TestHandlercase3(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	validated = true
	req := httptest.NewRequest(http.MethodGet, "/logout", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}
func TestHandlercase4(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	req := httptest.NewRequest(http.MethodGet, "/home", nil)
	w := httptest.NewRecorder()
	validated = true
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}

func TestHandlercase5(t *testing.T) { // needs db access
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	validated = true
	db.Connect("test")
	req := httptest.NewRequest(http.MethodGet, "/home/display", nil)
	w := httptest.NewRecorder()
	validated = true
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}
func TestHandlercase6(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/home/searchTitle", nil)
	w := httptest.NewRecorder()
	validated = true
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}

func TestHandlercase7(t *testing.T) { // needs db access
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	db.Loggers = Loggers
	req := httptest.NewRequest(http.MethodGet, "/home/searchTitle-Submit", nil)
	w := httptest.NewRecorder()
	validated = true
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}
func TestHandlercase8(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/home/searchUser", nil)
	w := httptest.NewRecorder()
	validated = true
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}

func TestHandlercase9(t *testing.T) { // needs db access
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	db.Loggers = Loggers
	req := httptest.NewRequest(http.MethodGet, "/home/searchUser-Submit", nil)
	w := httptest.NewRecorder()
	validated = true
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}
func TestHandlercase10(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/home/delete", nil)
	w := httptest.NewRecorder()
	validated = true
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}
func TestHandlercase11(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/home/delete-submit", nil)
	w := httptest.NewRecorder()
	validated = true
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}
func TestHandlercase12(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/home/create", nil)
	w := httptest.NewRecorder()
	validated = true
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}

func TestHandlercase13(t *testing.T) { // needs db access
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	db.Loggers = Loggers
	req := httptest.NewRequest(http.MethodGet, "/home/create-Submit", nil)
	w := httptest.NewRecorder()
	validated = true
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}
func TestHandlercase14(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/home/edit", nil)
	w := httptest.NewRecorder()
	validated = true
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}

func TestHandlercase15(t *testing.T) { // needs db access
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	db.Loggers = Loggers
	req := httptest.NewRequest(http.MethodGet, "/home/edit-submit", nil)
	w := httptest.NewRecorder()
	validated = true
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}
func TestHandlercase16(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	authtool.Loggers = Loggers
	req := httptest.NewRequest(http.MethodGet, "/createUser?username=test_username4&password=test_password1", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Error("expected http.StatusOk got ", w.Code)
	}
}

func TestHandlercase17(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/thisshouldntwork", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	if w.Code != http.StatusNotFound {
		t.Error("expected http.StatusNotFound got ", w.Code)
	}
	
	removeFiles()
}
