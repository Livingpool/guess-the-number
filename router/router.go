package router

import (
	"net/http"

	"github.com/Livingpool/constants"
	"github.com/Livingpool/handler"
	"github.com/Livingpool/views"
)

func Init() *http.ServeMux {
	router := http.NewServeMux()
	playerPool := handler.NewPlayerPool(constants.PLAYER_POOL_CAP)
	handler := handler.NewGameHandler(views.NewTemplates(), playerPool, &handler.RealTimeProvider{})

	// http.FS can be used to create a http.Filesystem
	var staticFS = http.FS(views.StaticFiles)
	fs := http.FileServer(staticFS)

	// Serve static files
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve all other requests
	router.HandleFunc("/", handler.Home)
	router.HandleFunc("GET /return", handler.ReturnHome)
	router.HandleFunc("POST /new", handler.NewGame)
	router.HandleFunc("GET /check", handler.CheckGuess)

	return router
}
