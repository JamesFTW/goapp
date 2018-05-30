package main

import (
  "fmt"
  "net/http"
  "database/sql"
  "github.com/gorilla/mux"
  _ "github.com/lib/pq"
)

func newRouter() *mux.Router {
  r := mux.NewRouter()

  r.HandleFunc("/hello", handler).Methods("GET")
  r.HandleFunc("/bird", getBirdHandler).Methods("GET")
  r.HandleFunc("/bird", createBirdHandler).Methods("POST")

  staticFileDirectory := http.Dir("./assets")
  staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

  r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

  return r
}
func main() {
  fmt.Println("Starting server...")
	connString := "postgres://jduphbzynbfuch:0dc67026b814a1156e0854d22222557848e675929fa4b72c8716a63863751bab@ec2-23-23-130-158.compute-1.amazonaws.com:5432/dcomiopimevuq7"
	db, err := sql.Open("postgres", connString)

	if err != nil {
		panic(err)
	}
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	InitStore(&dbStore{db: db})

	r := newRouter()
	fmt.Println("Serving on port 8080")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello World!")
}
