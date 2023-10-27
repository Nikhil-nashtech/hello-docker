package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	log.Println("Golang API")
	log.Println("Initializing Environment Variables")
	godotenv.Load()
	log.Println("Listening on localhost:4000")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println("GET for all courses")
	envVar := os.Getenv("DISPLAY_TEXT")
	w.Write([]byte(fmt.Sprintf("<h1>Hello %v</h1>", envVar)))
}
