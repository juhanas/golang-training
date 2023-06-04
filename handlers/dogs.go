package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

var dogsFilePath = "./data/dogs.txt"
var dogs []string

// getDog returns the requested dog
// Returns an error if the dog is not found
func getDog(dog string, dogs []string) (string, error) {
	for _, v := range dogs {
		if strings.Contains(v, dog) {
			return v, nil
		}
	}
	return "", fmt.Errorf("could not find dog with name '%v'", dog)
}

// GetDogs reads dogs from the file and inserts them into the memory
func GetDogs() {
	dogNames := []string{
		"Charlie", "Buddy", "Cooper",
	}

	// Todo: Parse names from file

	dogs = dogNames
}

// GetDogsHandler returns all dogs
func GetDogsHandler(w http.ResponseWriter, r *http.Request) {
	// Todo: Return data as json
	io.WriteString(w, fmt.Sprint(dogs))
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

	io.WriteString(w, "Your chosen dog: "+dog)
}
