package handler

import (
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/Livingpool/views"
)

type GameHandler struct {
	renderer *views.Templates
}

func NewGameHandler(renderer *views.Templates) *GameHandler {
	return &GameHandler{
		renderer: renderer,
	}
}

func (h *GameHandler) Home(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, "home", nil)
}

func (h *GameHandler) NewGame(w http.ResponseWriter, r *http.Request) {
	digit, err := strconv.Atoi(r.FormValue("digit"))
	if err != nil {
		http.Error(w, "Invalid input", http.StatusUnprocessableEntity)
		return
	}
	if digit < 1 || digit > 10 {
		http.Error(w, "This digit is currently not supported", http.StatusUnprocessableEntity)
		return
	}

	answer := rand.Intn(int(math.Pow(10, float64(digit-1))))
	answerStr := strconv.Itoa(answer)
	for i := 0; i < digit-len(answerStr); i++ {
		answerStr = "0" + answerStr
	}

	// Execute the template
	h.renderer.Render(w, "game", nil)
}

// func (h GameHandler) CheckGuess(w http.ResponseWriter, r *http.Request) string {
// 	a := 0
// 	b := 0
// 	for i := 0; i < len(guess); i++ {
// 		if guess[i] == answer[i] {
// 			a++
// 		} else if strings.Contains(answer, string(guess[i])) {
// 			b++
// 		}
// 	}

// 	countA := strconv.Itoa(a)
// 	countB := strconv.Itoa(b)

// 	return countA + "a" + countB + "b"
// }
