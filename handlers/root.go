package handlers

import (
	"io"
	"net/http"
)

// GetRootHandler returns general message and api
func GetRootHandler(w http.ResponseWriter, r *http.Request) {
	message := `Welcome to Pet store!
This store contains several pets that you can get details of.
Eventually you can also add new pets to the store and update details for existing ones.

Available paths:
GET /cats
GET /cat/ - use queryParam "name"`
	io.WriteString(w, message)
}
