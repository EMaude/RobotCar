package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test Successful"))
}

func home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/static/", 301)
}

func main() {
	//routes
	router := mux.NewRouter()

	router.HandleFunc("/api/test", test).Methods("GET")

	router.HandleFunc("/", home)
	//Static Files
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))

	log.Println("Listeing on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
