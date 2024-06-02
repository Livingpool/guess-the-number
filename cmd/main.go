package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Livingpool/middleware"
	"github.com/Livingpool/router"
)

func main() {
	router := router.Init()

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "42069"
		log.Printf("Defaulting to port %s\n", port)
	}
	port = ":" + port

	server := http.Server{
		Addr:    port,
		Handler: stack(router),
	}

	fmt.Println("Server listening on port", port)
	log.Fatal(server.ListenAndServe())
}
