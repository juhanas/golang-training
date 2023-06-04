package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/juhanas/golang-training/handlers"
	"github.com/julienschmidt/httprouter"
)

// mustRunServer starts the server or exits if an error is encountered
// The server can be accessed at 127.0.0.1:3333
func mustRunServer(router *httprouter.Router) {
	fmt.Println("Starting server at 127.0.0.1:3333")

	// Initialize the dogs
	handlers.GetDogs()

	err := http.ListenAndServe("127.0.0.1:3333", router)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
