package handlers

import (
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/julienschmidt/httprouter"
)

var cats = map[string]string{
	"a": "Alice",
	"b": "Bella",
	"c": "Coco",
}

// sortKeys returns a list of the keys in the given map, sorted alphabetically
func sortKeys(data map[string]string) []string {
	keys := []string{}
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// getCatList returns a list of all cats
func getCatList(cats map[string]string) []string {
	keys := sortKeys(cats)

	catList := []string{}
	for _, key := range keys {
		catList = append(catList, cats[key])
	}
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

// findNameKey returns the key for the given name
// Returns error if name is reserved or key could not be found
func findNameKey(name string, dict map[string]string, i int) string {
	key := name[:i]
	_, ok := dict[key]
	if !ok {
		return key
	}
	return findNameKey(name, dict, i+1)
}

// PostCatHandler adds a new cat
func PostCatHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	name := r.Form["name"][0]
	key := findNameKey(name, cats, 1)
	cats[key] = name

	w.WriteHeader(http.StatusOK)
}
