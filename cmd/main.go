package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/Livingpool/middleware"
	"github.com/Livingpool/router"
)

func main() {
	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{})
	logger := slog.New(h)

	config := middleware.LoggingConfig{
		DefaultLevel:     slog.LevelInfo,
		ServerErrorLevel: slog.LevelError,
		ClientErrorLevel: slog.LevelWarn,
	}

	stack := middleware.CreateStack(
		middleware.Logging(logger, config),
	)

	router := router.Init()

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
