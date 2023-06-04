package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	utils "github.com/juhanas/golang-training/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDogs(t *testing.T) {
	// Allow locating the txt-file when running tests
	dogsFilePath = "../data/dogs.txt"

	expected := []string{
		"Charlie",
		"Buddy",
		"Cooper",
		"Duke",
		"Max",
		"Rocky",
		"Stella",
		"Molly",
		"Ace",
		"Milo",
		"Sadie",
	}

	GetDogs()
	ok := utils.CompareLists(dogs, expected)
	require.True(t, ok, "lists do not match", dogs, expected)
}

func TestGetDogsHandler(t *testing.T) {
	dogs = []string{
		"Charlie",
		"Buddy",
		"Cooper",
	}

	req, err := http.NewRequest("GET", "/dogs", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetDogsHandler)

	handler.ServeHTTP(rr, req)

	wantedStatus := http.StatusOK
	assert.Equal(t, rr.Code, wantedStatus, "handler returned wrong status")

	dogsReceived := []string{}
	json.Unmarshal(rr.Body.Bytes(), &dogsReceived)
	expected := []string{"Charlie", "Buddy", "Cooper"}

	assert.Equal(t, expected, dogsReceived)
}

func TestGetDog(t *testing.T) {
	dogs = []string{
		"Charlie",
		"Buddy",
		"Cooper",
	}

	router := httprouter.New()
	router.GET("/dog/:name", GetDogHandler)

	req, err := http.NewRequest("GET", "/dog/Coo", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	wantedStatus := http.StatusOK
	require.Equal(t, rr.Code, wantedStatus)

	expected := "Your chosen dog: Cooper"
	assert.Equal(t, rr.Body.String(), expected)
}

func TestGetDogNotFound(t *testing.T) {
	dogs = []string{}

	router := httprouter.New()
	router.GET("/dog/:name", GetDogHandler)

	req, err := http.NewRequest("GET", "/dog/not-found", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	wantedStatus := http.StatusBadRequest
	assert.Equal(t, rr.Code, wantedStatus)

	expected := "could not find dog with name 'not-found'"
	assert.Equal(t, rr.Body.String(), expected)
}
