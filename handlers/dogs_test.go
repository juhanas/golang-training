package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/juhanas/golang-training/models"
	utils "github.com/juhanas/golang-training/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDogs(t *testing.T) {
	// Allow locating the txt-file when running tests
	dogsFilePath = "../data/dogs.txt"

	expected := []models.Dog{
		utils.CreateDog("Charlie"),
		utils.CreateDog("Buddy"),
		utils.CreateDog("Cooper"),
		utils.CreateDog("Duke"),
		utils.CreateDog("Max"),
		utils.CreateDog("Rocky"),
		utils.CreateDog("Stella"),
		utils.CreateDog("Molly"),
		utils.CreateDog("Ace"),
		utils.CreateDog("Milo"),
		utils.CreateDog("Sadie"),
	}

	GetDogs()
	ok := utils.CompareDogs(dogs, expected)
	require.True(t, ok, "lists do not match", dogs, expected)
}

func TestGetDogsHandler(t *testing.T) {
	dogs = []models.Dog{
		utils.CreateDog("Charlie"),
		utils.CreateDog("Buddy"),
		utils.CreateDog("Cooper"),
	}

	req, err := http.NewRequest("GET", "/dogs", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetDogsHandler)

	handler.ServeHTTP(rr, req)

	wantedStatus := http.StatusOK
	assert.Equal(t, rr.Code, wantedStatus, "handler returned wrong status")

	dogsReceived := []models.Dog{}
	json.Unmarshal(rr.Body.Bytes(), &dogsReceived)

	ok := utils.CompareDogs(dogsReceived, dogs)
	assert.True(t, ok, "dog lists differ", dogs, dogsReceived)
}

func TestGetDog(t *testing.T) {
	dogs = []models.Dog{
		utils.CreateDog("Charlie"),
		utils.CreateDog("Buddy"),
		utils.CreateDog("Cooper"),
	}

	router := httprouter.New()
	router.GET("/dog/:name", GetDogHandler)

	req, err := http.NewRequest("GET", "/dog/Coo", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	wantedStatus := http.StatusOK
	require.Equal(t, rr.Code, wantedStatus)

	expected := "Your chosen dog: Cooper whose color is black and pack: Who let the dogs out?"
	assert.Equal(t, rr.Body.String(), expected)
}

func TestGetDogNotFound(t *testing.T) {
	dogs = []models.Dog{}

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
