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
	answer   string
}

func NewGameHandler(renderer *views.Templates) *GameHandler {
	return &GameHandler{
		renderer: renderer,
	}
}

type FormData struct {
	Digit int
	Error string
}

func (h *GameHandler) Home(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, "base", nil)
}

func (h *GameHandler) NewGame(w http.ResponseWriter, r *http.Request) {
	digit, err := strconv.Atoi(r.FormValue("digit"))
	if err != nil {
		formData := FormData{
			Digit: 0,
			Error: "Input is not a digit :(",
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.renderer.Render(w, "form", formData)
		return
	}
	if digit < 1 || digit > 10 {
		formData := FormData{
			Digit: 0,
			Error: "Input is not in range :(",
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.renderer.Render(w, "form", formData)
		return
	}

	answer := rand.Intn(int(math.Pow(10, float64(digit-1))))
	answerStr := strconv.Itoa(answer)
	for i := 0; i < digit-len(answerStr); i++ {
		answerStr = "0" + answerStr
	}

	h.answer = answerStr

	formData := FormData{
		Digit: digit,
		Error: "",
	}
	// Execute the template
	h.renderer.Render(w, "game", formData)
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
