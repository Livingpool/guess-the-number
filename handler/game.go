package handler

import (
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Livingpool/views"
)

type GameHandler struct {
	renderer *views.Templates
	answer   string
	results  GuessResults
}

func NewGameHandler(renderer *views.Templates) *GameHandler {
	return &GameHandler{
		renderer: renderer,
	}
}

type FormData struct {
	Digit int
	Start string
	End   string
	Error string
}

type ResultRow struct {
	TimeStamp string
	Guess     string
	Result    string
}

type GuessResults struct {
	Rows []ResultRow
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
	for len(answerStr) < digit {
		answerStr = "0" + answerStr
	}

	h.answer = answerStr

	start, end := "", ""
	for range digit {
		start += "0"
		end += "9"
	}

	formData := FormData{
		Digit: digit,
		Start: start,
		End:   end,
		Error: "",
	}
	// Execute the template
	h.renderer.Render(w, "game", formData)
}

func (h *GameHandler) CheckGuess(w http.ResponseWriter, r *http.Request) {
	guessStr := r.URL.Query().Get("guess")
	// Handle error cases
	if len(guessStr) != len(h.answer) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		row := ResultRow{
			TimeStamp: time.Now().Format(time.DateTime),
			Guess:     guessStr,
			Result:    "invalid input len :(",
		}
		h.results.Rows = append([]ResultRow{row}, h.results.Rows...)
		h.renderer.Render(w, "result", h.results)
		return
	}

	a, b := 0, 0
	aMap := make([]bool, len(guessStr)) // positions of a's
	countMap := make([]int, 10)         // occurences of chars (for calc b's)

	for i := range 10 {
		countMap[i] = strings.Count(h.answer, strconv.Itoa(i))
	}

	log.Println(guessStr, h.answer)
	for i := 0; i < len(guessStr); i++ {
		c, _ := strconv.Atoi(string(guessStr[i]))
		if guessStr[i] == h.answer[i] {
			if countMap[c] <= 0 {
				b--
				a++
			} else {
				a++
				countMap[c]--
			}
			aMap[i] = true
		} else {
			if countMap[c] > 0 {
				b++
				countMap[c]--
			}
		}
	}

	countA := strconv.Itoa(a)
	countB := strconv.Itoa(b)
	result := countA + "a" + countB + "b"

	row := ResultRow{
		TimeStamp: time.Now().Format(time.DateTime),
		Guess:     guessStr,
		Result:    result,
	}

	h.results.Rows = append([]ResultRow{row}, h.results.Rows...)

	h.renderer.Render(w, "result", h.results)
}
