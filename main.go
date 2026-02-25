package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(hndler),
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
