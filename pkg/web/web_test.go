package web

import (
	"golock3r/server/db"
	"golock3r/server/logger"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestLandingPage(t *testing.T) {
	os.Chdir("../")
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	var landingpage = login(w, req)
	if !landingpage {
		t.Error("expected to have the landing page shown landing page didnt show ")
	}
}
func TestLoginSubmit(t *testing.T){
	Loggers = logger.CreateLoggers("testlogs.txt")
	os.Chdir("../")
	w := httptest.NewRecorder()
	form := url.Values{} 
	form.Add("username", "demo")
	form.Add("password", "demo123")
	req := httptest.NewRequest(http.MethodPost, "/login-submit",strings.NewReader(form.Encode()))
	req.Form = form 
	var loginSubmit = loginSubmit(w,req)
	if !loginSubmit{
		t.Error("login should have submitted succesfully it didnt")
	}

}
func TestCreateUser(t *testing.T) {
	os.Chdir("../")
	Loggers = logger.CreateLoggers("testlogs.txt")
	w := httptest.NewRecorder()
	form := url.Values{} 
	form.Add("username", "demo")
	form.Add("password", "demo123")
	req := httptest.NewRequest(http.MethodGet, "/createUser",strings.NewReader(form.Encode()))
	req.Form = form 
	var createUser = createUser(w,req)
	if !createUser{
		t.Error("this didnt work")
	}
	

	
}
func TestLogout(t *testing.T) {
	db.Connect("demo")
	
	os.Chdir("../")
	req := httptest.NewRequest(http.MethodGet, "/logout", nil)
	w := httptest.NewRecorder()
	var logout = logout(w,req)
	if !logout {
		t.Error("logout unseccessful")
	}
}
func TestReadall(t *testing.T) {
	Loggers = logger.CreateLoggers("testlogs.txt")
	db.Connect("demo")
	os.Chdir("../")
	req := httptest.NewRequest(http.MethodGet, "/home/display", nil)
	w := httptest.NewRecorder()
	var readAll = readAll(w,req)
	if !readAll{
		t.Error("unable to read all")

	}
}
func TestSearchByTitle(t *testing.T){
	os.Chdir("../")
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/searchTitle", nil)
	w := httptest.NewRecorder()
	var landingpage = searchByTitle(w, req)
	if !landingpage {
		t.Error("expected to have the search by title page shown")
	}

}
func TestSearchByTitleSubmit(t *testing.T){

}
func TestSearchByUsername(t *testing.T){
	os.Chdir("../")
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/searchUser", nil)
	w := httptest.NewRecorder()
	var landingpage = searchByUsername(w, req)
	if !landingpage {
		t.Error("expected to have the search by username page shown")
	}
}
func TestSearchByUsernameSubmit(t *testing.T){

}

func TestDelete(t *testing.T){
	os.Chdir("../")
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/delete", nil)
	w := httptest.NewRecorder()
	var landingpage = login(w, req)
	if !landingpage {
		t.Error("expected to have the delete page shown")
	}
}

func TestDeleteSubmit(t *testing.T){

}

func TestCreateEntry(t *testing.T){
	os.Chdir("../")
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/create", nil)
	w := httptest.NewRecorder()
	var landingpage = login(w, req)
	if !landingpage {
		t.Error("expected to have the create entry page shown")
	}
}

func TestCreateEntrySubmit(t* testing.T){

}
func TestEdit(t *testing.T){
	os.Chdir("../")
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home/edit", nil)
	w := httptest.NewRecorder()
	var landingpage = login(w, req)
	if !landingpage {
		t.Error("expected to have the edit page didnt show ")
	}
}

func TestEditSubmit(t *testing.T){

}

func TestHome(t *testing.T){
	os.Chdir("../")
	Loggers = logger.CreateLoggers("testlogs.txt")
	req := httptest.NewRequest(http.MethodGet, "/home", nil)
	w := httptest.NewRecorder()
	var landingpage = home(w, req)
	if !landingpage {
		t.Error("expected to have the home page shown landing page didnt show ")
	}
}