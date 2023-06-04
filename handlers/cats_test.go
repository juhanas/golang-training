package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	utils "github.com/juhanas/golang-training/utils"
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

func runGetCatListTest(t *testing.T, testName string, cats map[string]string, expected []string) {
	finalCats := getCatList(cats)
	ok := utils.CompareLists(finalCats, expected)
	if !ok {
		t.Errorf("Unexpected list received for test %s. Want: %v - received: %v", testName, expected, finalCats)
	}
}

func TestGetCatList(t *testing.T) {
	testCases := []struct {
		name     string
		input    map[string]string
		expected []string
	}{
		{
			"success",
			map[string]string{
				"a": "Alice",
				"b": "Bella",
				"c": "Coco",
			},
			[]string{
				"Alice",
				"Bella",
				"Coco",
			},
		},
		{
			"emptyDict",
			map[string]string{},
			[]string{},
		},
		{
			"emptyName",
			map[string]string{
				"a": "",
				"b": "Bella",
				"c": "Coco",
			},
			[]string{
				"",
				"Bella",
				"Coco",
			},
		},
	}
	for _, tc := range testCases {
		runGetCatListTest(t, tc.name, tc.input, tc.expected)
	}
}

func TestPostCat(t *testing.T) {
	originalCats := utils.CopyMap(cats)
	// Defer function is executed when this function exits - no matter for what reason
	defer func() {
		cats = originalCats
	}()

	router := httprouter.New()
	router.POST("/cat/", PostCatHandler)

	data := url.Values{}
	data.Set("name", "accident")
	body := strings.NewReader(data.Encode())

	req, err := http.NewRequest("POST", "/cat/", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	wantedStatus := http.StatusOK
	if status := rr.Code; status != wantedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, wantedStatus)
	}

	// Verify the data struct was changed
	expectedCats := []string{"Alice", "accident", "Bella", "Coco"}
	actualCats := getCatList(cats)
	if ok := utils.CompareLists(actualCats, expectedCats); !ok {
		t.Errorf("unexpected cats: got %v want %v",
			actualCats, expectedCats)
	}
}
