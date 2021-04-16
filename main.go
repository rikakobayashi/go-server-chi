package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Test struct {
	Message string `json:"message"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	msg := "hello"
	t := Test{Message: msg}
	e := json.NewEncoder(w)
	if err := e.Encode(t); err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/ping", testHandler)
	http.ListenAndServe(":3000", r)
}