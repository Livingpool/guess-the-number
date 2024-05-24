package router

import (
	"net/http"

	"github.com/Livingpool/handler"
	"github.com/Livingpool/views"
)

func Init() *http.ServeMux {
	router := http.NewServeMux()
	handler := handler.NewGameHandler(views.NewTemplates())

	css := http.FileServer(http.Dir("./views/css"))
	router.Handle("/styles/", http.StripPrefix("/styles/", css))

	router.HandleFunc("GET /home", handler.Home)
	router.HandleFunc("POST /new", handler.NewGame)
	router.HandleFunc("GET /check", handler.CheckGuess)

	return router
}
