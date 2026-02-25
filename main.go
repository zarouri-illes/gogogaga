package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/api", hndler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Failed", err)
	}

	fmt.Println("test")
}

func hndler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ping"))
}
