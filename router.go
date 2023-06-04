package main

import (
	"github.com/juhanas/golang-training/handlers"
	"github.com/julienschmidt/httprouter"
)

// createRouter creates a new router and sets all routes for the server
func createRouter() *httprouter.Router {
	router := httprouter.New()
	router.HandlerFunc("GET", "/", handlers.GetRootHandler)
	router.GET("/cat/", handlers.GetCatHandler)
	router.HandlerFunc("GET", "/cats", handlers.GetCatsHandler)
	return router
}
