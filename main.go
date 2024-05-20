/*
 * GSES BTC application
 * API version: 1.0.0
 */

package main

import (
	exchangerate "exchangerate/go"
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
	"log"
	"net/http"
)

func main() {

	exchangerate.DbInit()

	router := exchangerate.NewRouter()
	log.Printf("Server started")

	log.Fatal(http.ListenAndServe(":8080", router))
}
