package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func TestReplaceDog(t *testing.T) {
	dogs = []models.Dog{
		{
			Animal: models.Animal{
				Name:  "dog",
				Color: "black",
			},
			Pack: "pack",
		},
	}
	newDog := models.Dog{
		Animal: models.Animal{
			Name:  "dog",
			Color: "red",
		},
		Pack: "newPack",
	}

	err := replaceDog(&newDog)
	require.NoError(t, err)
	require.Equal(t, 1, len(dogs), fmt.Sprintf("Wrong amount of dogs stored: got %v want %v", len(dogs), 1))

	assert.Equal(t, newDog.Name, dogs[0].Name)
	assert.Equal(t, newDog.Color, dogs[0].Color)
	assert.Equal(t, newDog.Pack, dogs[0].Pack)
}

func TestUpdateDog(t *testing.T) {
	dogs = []models.Dog{
		{
			Animal: models.Animal{
				Name:  "dog",
				Color: "black",
			},
			Pack: "pack",
		},
	}

	router := httprouter.New()
	router.POST("/dog/", UpdateDogHandler)

	var jsonData = []byte(`{
		"name": "dog",
		"color": "red",
		"pack": "newPack"
	}`)

	req, err := http.NewRequest("POST", "/dog/", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	wantedStatus := http.StatusOK
	if status := rr.Code; status != wantedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, wantedStatus)
	}

	expected := "Dog dog color changed to red and pack to newPack"
	assert.Equal(t, rr.Body.String(), expected)

	// Verify the data struct was changed
	expectedDogs := []models.Dog{{Animal: models.Animal{Name: "dog", Color: "red"}, Pack: "newPack"}}
	ok := utils.CompareDogs(dogs, expectedDogs)
	assert.True(t, ok, "dog lists differ", dogs, expectedDogs)
}
