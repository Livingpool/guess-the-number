package handler

import (
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

type GameHandler struct{}

func (h *GameHandler) NewGame(w http.ResponseWriter, r *http.Request) {
	digits, err := strconv.Atoi(r.URL.Query().Get("digits"))
	if err != nil {
		http.Error(w, "Invalid input", http.StatusUnprocessableEntity)
		return
	}
	if digits < 1 || digits > 10 {
		http.Error(w, "This digit is currently not supported", http.StatusUnprocessableEntity)
		return
	}

	answer := rand.Intn(int(math.Pow(10, float64(digits-1))))
	answerStr := strconv.Itoa(answer)
	for i := 0; i < digits-len(answerStr); i++ {
		answerStr = "0" + answerStr
	}

	// Execute the template
}

func (h GameHandler) CheckGuess(w http.ResponseWriter, r *http.Request) string {
	a := 0
	b := 0
	for i := 0; i < len(guess); i++ {
		if guess[i] == answer[i] {
			a++
		} else if strings.Contains(answer, string(guess[i])) {
			b++
		}
	}

	countA := strconv.Itoa(a)
	countB := strconv.Itoa(b)

	return countA + "a" + countB + "b"
}
