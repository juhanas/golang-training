package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var cats = map[string]string{
	"a": "Alice",
	"b": "Bella",
	"c": "Coco",
}

// getCatList returns a list of all cats
func getCatList(cats map[string]string) []string {
	catList := []string{}
	return catList
}

// GetCatsHandler returns all cats
func GetCatsHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprint(getCatList(cats)))
}

// GetCatHandler returns details of a single cat, by name
func GetCatHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := r.URL.Query().Get("name")

	cat, ok := cats[name]
	if !ok || cat == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "could not find cat with name '"+name+"'")
		return
	}

	io.WriteString(w, "Your chosen cat: "+cat)
}
