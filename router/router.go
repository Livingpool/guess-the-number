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
	handler := handler.NewGameHandler(views.NewTemplates(), playerPool)

	css := http.FileServer(http.Dir("./views/css"))
	router.Handle("/styles/", http.StripPrefix("/styles/", css))

	assets := http.FileServer(http.Dir("./views/assets"))
	router.Handle("/assets/", http.StripPrefix("/assets/", assets))

	scripts := http.FileServer(http.Dir("./views/scripts"))
	router.Handle("/scripts/", http.StripPrefix("/scripts/", scripts))

	router.HandleFunc("GET /home", handler.Home)
	router.HandleFunc("GET /return", handler.ReturnHome)
	router.HandleFunc("POST /new", handler.NewGame)
	router.HandleFunc("GET /check", handler.CheckGuess)

	return router
}
