package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Livingpool/middleware"
	"github.com/Livingpool/router"
)

func main() {
	router := router.Init()

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":42069",
		Handler: stack(router),
	}

	fmt.Println("Server listening on port 42069")
	log.Fatal(server.ListenAndServe())
}
