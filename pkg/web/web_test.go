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
		t.Error("expected to have the landing page shown landing page didnt show ")
	}
}
func TestLoginSubmit(t *testing.T){
	os.Chdir("../")
	w := httptest.NewRecorder()
	form := url.Values{} 
	form.Add("username", "demo")
	form.Add("password", "demo123")
	req := httptest.NewRequest("POST", "/",strings.NewReader(form.Encode()))
	req.Form = form 
	var loginSubmit = loginSubmit(w,req)
	if !loginSubmit{
		t.Error("login should have submitted succesfully it didnt")
	}

}
func TestCreateUser(t *testing.T) {
	os.Chdir("../")
	
}