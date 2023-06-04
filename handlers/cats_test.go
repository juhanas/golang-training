package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestGetCats(t *testing.T) {
	req, err := http.NewRequest("GET", "/cats", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	// Use golang's net/http handler
	handler := http.HandlerFunc(GetCatsHandler)

	handler.ServeHTTP(rr, req)

	wantedStatus := http.StatusOK
	if status := rr.Code; status != wantedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, wantedStatus)
	}

	expected := `[Alice Bella Coco]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetCat(t *testing.T) {
	req, err := http.NewRequest("GET", "/cat/?name=a", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	// Use external httprouter package's handler to allow path params
	router := httprouter.New()
	router.GET("/cat/", GetCatHandler)

	router.ServeHTTP(rr, req)

	wantedStatus := http.StatusOK
	if status := rr.Code; status != wantedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, wantedStatus)
	}

	expected := "Your chosen cat: Alice"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetCatNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/cat/?name=not-found", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	// Use external httprouter package's handler to allow path params
	router := httprouter.New()
	router.GET("/cat/", GetCatHandler)

	router.ServeHTTP(rr, req)

	wantedStatus := http.StatusBadRequest
	if status := rr.Code; status != wantedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, wantedStatus)
	}

	expected := "could not find cat with name 'not-found'"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
