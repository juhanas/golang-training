package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/juhanas/golang-training/models"
	"github.com/juhanas/golang-training/utils"
	"github.com/julienschmidt/httprouter"
)

var dogsFilePath = "./data/dogs.txt"
var dogs []models.Dog

// getDog returns the requested dog
// Returns an error if the dog is not found
func getDog(dog string, dogs []models.Dog) (*models.Dog, error) {
	for _, v := range dogs {
		if strings.Contains(v.Name, dog) {
			return &v, nil
		}
	}
	return nil, fmt.Errorf("could not find dog with name '%v'", dog)
}

// GetDogs reads dogs from the file and inserts them into the memory
func GetDogs() {
	data, err := os.ReadFile(dogsFilePath)
	if err != nil {
		panic(err)
	}
	dogNames := utils.ParseCSVDataToArray(data)

	dogsMap := []models.Dog{}
	for _, dogName := range dogNames {
		dog := models.Dog{
			Animal: models.Animal{
				Name:  dogName,
				Color: "black",
			},
			Pack: "Who let the dogs out?",
		}
		dogsMap = append(dogsMap, dog)
	}

	dogs = dogsMap
}

// GetDogsHandler returns all dogs
func GetDogsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dogs)
}

// GetDogHandler returns details of a single dog requested by name
// Returns an error if the dog is not found
func GetDogHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := ps.ByName("name")

	dog, err := getDog(name, dogs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, "Your chosen dog: "+dog.Name+" whose color is "+dog.Color+" and pack: "+dog.Pack)
}
