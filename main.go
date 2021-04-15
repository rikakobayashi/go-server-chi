package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Test struct {
	Message string `json:"message"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	message := "hello"
	test := Test{Message: message}
	data, _ := json.Marshal(test)
	w.Write(data)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", testHandler)
	http.ListenAndServe(":3000", r)
}