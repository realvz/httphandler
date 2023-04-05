package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	CONN_HOST = ""
	CONN_PORT = "8080"
)

var GetRequestHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Generate a random number between 0 and 99
	randomNum := rand.Intn(70)

	if randomNum < 10 {
		// Return a 500 Internal Server Error response
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else if randomNum < 20 {
		// Return a 400 Bad Request response
		http.Error(w, "Bad Request", http.StatusBadRequest)
	} else {
		// Return a normal response
		w.Write([]byte("Hello World!"))
	}
})

var PostRequestHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Generate a random number between 0 and 99
	randomNum := rand.Intn(100)

	if randomNum < 10 {
		// Return a 500 Internal Server Error response
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} else if randomNum < 20 {
		// Return a 400 Bad Request response
		http.Error(w, "Bad Request", http.StatusBadRequest)
	} else {
		// Return a normal response
		w.Write([]byte("It's a Post Request!"))
	}
})

var PathVariableHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	w.Write([]byte("Hi " + name))
})

func main() {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().Unix())

	router := mux.NewRouter()

	router.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(GetRequestHandler))).Methods("GET")

	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error starting http server : ", err)
	}
	defer logFile.Close()

	router.Handle("/post", handlers.LoggingHandler(logFile, PostRequestHandler)).Methods("POST")

	router.Handle("/hello/{name}", handlers.CombinedLoggingHandler(logFile, PathVariableHandler)).Methods("GET")

	log.Printf("Listening on %s:%s...\n", CONN_HOST, CONN_PORT)
	err = http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server : ", err)
	}
}
