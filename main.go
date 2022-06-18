package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/N-Iwata/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	fmt.Println("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
