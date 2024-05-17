package router

import (
	"net/http"
	"path"

	"github.com/Livingpool/handler"
	"github.com/Livingpool/web"
)

type Count struct {
	Count int
}

func Init() *http.ServeMux {
	router := http.NewServeMux()
	handler := &handler.GameHandler{}

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	count := Count{Count: 0}
	router.HandleFunc("GET /count", func(w http.ResponseWriter, r *http.Request) {
		count.Count++
		fp := path.Join("web", "home.html")
		if err := web.Render(w, fp, count); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	router.HandleFunc("POST /newgame", handler.NewGame)

	return router
}
