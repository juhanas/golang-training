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
	router.POST("/cat/", handlers.PostCatHandler)
	router.HandlerFunc("GET", "/cats", handlers.GetCatsHandler)
	router.GET("/dog/:name", handlers.GetDogHandler)
	router.HandlerFunc("GET", "/dogs", handlers.GetDogsHandler)
	return router
}
