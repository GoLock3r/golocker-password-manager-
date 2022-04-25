package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestLandingPage(t *testing.T) {
	os.Chdir("../")
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	var landingpage = login(w, req)
	if !landingpage {
		t.Error("Landing page not successfully displayed.")
	}
}

func TestLoginSubmit(t *testing.T){
	os.Chdir("../")
	w := httptest.NewRecorder()
	form := url.Values{} 
	form.Add("username", "demo")
	form.Add("password", "demo123")
	req := httptest.NewRequest(http.MethodPost, "/login-submit",strings.NewReader(form.Encode()))
	req.Form = form 
	var loginSubmit = loginSubmit(w,req)
	if !loginSubmit{
		t.Error("Login unsuccessful.")
	}
}

func TestCreateUser(t *testing.T) {
	os.Chdir("../")
	w := httptest.NewRecorder()
	form := url.Values{} 
	form.Add("username", "demo")
	form.Add("password", "demo123")
	req := httptest.NewRequest(http.MethodPost, "/createUser",strings.NewReader(form.Encode()))
	req.Form = form 
	var createUser = createUser(w,req)
	if !createUser{
		t.Error("User not created.")
	}	
}