package web

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
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
	req:= httptest.NewRequest(http.MethodGet, "/login-submit?username=demo&password=demo123",nil)
	w:= httptest.NewRecorder()
	var loginSubmit = loginSubmit(w,req)
	if !loginSubmit{
		t.Error("login should have submitted succesfully it didnt")
	}

}