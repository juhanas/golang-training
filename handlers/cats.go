package handlers

import (
	"fmt"
	"io"
	"net/http"

	utils "github.com/juhanas/golang-training/utils"
	"github.com/julienschmidt/httprouter"
)

var cats = map[string]string{
	"a": "Alice",
	"b": "Bella",
	"c": "Coco",
}

// getCatList returns a list of all cats
func getCatList(cats map[string]string) []string {
	keys := utils.SortKeys(cats)

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

// PostCatHandler adds a new cat
func PostCatHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	if r.Form == nil || r.Form["name"] == nil || len(r.Form["name"]) < 1 || r.Form["name"][0] == "" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "no name for cat given")
		return
	}

	name := r.Form["name"][0]
	key, err := utils.FindNameKey(name, cats, 1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	cats[key] = name
	w.WriteHeader(http.StatusOK)
}
