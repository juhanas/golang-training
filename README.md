# golang-training

## Description
This repository contains example code for learning Golang.

## Installation
1. Clone the repo
2. Run `go mod tidy` to download and install dependencies

## Usage
1. Start the server with `go run .`
2. Make GET request to `127.0.0.1:3333/` to see the API
    - Use for example `curl http://localhost:3333`
3. To make POST requests you can use the following command
    - `curl -d "name=accident" -X POST http://127.0.0.1:3333/cat/`
4. To make PUT requests with json you can use the following command
    - `curl -d '{"name": "Charlie","color":"red","pack":"new"}' -H "Content-type:application/json" -X PUT http://127.0.0.1:3333/dog`

## Testing
To run all tests, use `go test ./...`
To run a single test, use `go test -run TestName ./...`