package main

func main() {
	router := createRouter()
	mustRunServer(router)
}
