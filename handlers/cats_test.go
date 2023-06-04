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

// compareLists returns if the two lists contain the same data
func compareLists(list1, list2 []string) bool {
	if list1 == nil && list2 != nil || list1 != nil && list2 == nil {
		return false
	}
	if len(list1) != len(list2) {
		return false
	}
	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

func runGetCatListTest(t *testing.T, testName string, cats map[string]string, expected []string) {
	finalCats := getCatList(cats)
	ok := compareLists(finalCats, expected)
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
