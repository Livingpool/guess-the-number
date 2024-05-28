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
	renderer   *views.Templates
	playerPool *PlayerPool
}

func NewGameHandler(renderer *views.Templates, playerPool *PlayerPool) *GameHandler {
	return &GameHandler{
		renderer:   renderer,
		playerPool: playerPool,
	}
}

type FormData struct {
	Digit    int
	Start    string
	End      string
	Error    string
	PlayerId int
}

type resultRow struct {
	TimeStamp string
	Guess     string
	Result    string
}

type guessResults struct {
	Rows []resultRow
}

func (h *GameHandler) Home(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, "base", nil)
}

func (h *GameHandler) ReturnHome(w http.ResponseWriter, r *http.Request) {
	h.renderer.Render(w, "home", nil)
}

func (h *GameHandler) NewGame(w http.ResponseWriter, r *http.Request) {
	digit, err := strconv.Atoi(r.FormValue("digit"))

	// Invalid input error
	if err != nil {
		formData := FormData{
			Error: "Input is not a digit :(",
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.renderer.Render(w, "form", formData)
		return
	}

	// Input value not in range error
	if digit < 1 || digit > 10 {
		formData := FormData{
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

	// PlayerPool full error
	newPlayer := NewPlayer(answerStr)
	if err = h.playerPool.AddPlayer(newPlayer); err != nil {
		formData := FormData{
			Error: "Server is full. Please try again later!",
		}
		log.Println(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.renderer.Render(w, "form ", formData)
		return
	}

	start, end := "", ""
	for range digit {
		start += "0"
		end += "9"
	}

	formData := FormData{
		Digit:    digit,
		Start:    start,
		End:      end,
		Error:    "",
		PlayerId: newPlayer.Id,
	}

	log.Println("Player registered", newPlayer.Id)

	// Execute the template
	h.renderer.Render(w, "game", formData)
}

// TODO: handle player id errors
// TODO: set idle timeouts

func (h *GameHandler) CheckGuess(w http.ResponseWriter, r *http.Request) {
	guessStr := r.URL.Query().Get("guess")
	playerId := r.URL.Query().Get("id")

	// Player id not parseable error
	id, err := strconv.Atoi(playerId)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Player doesn't exist error
	player, exists := h.playerPool.players[id]
	if !exists {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// Guess/Answer length not match error
	if len(guessStr) != len(player.Answer) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		row := resultRow{
			TimeStamp: time.Now().Format(time.DateTime),
			Guess:     guessStr,
			Result:    "invalid input len :(",
		}
		player.GuessResults.Rows = append([]resultRow{row}, player.GuessResults.Rows...)
		h.renderer.Render(w, "result", player.GuessResults.Rows)
		return
	}

	a, b := 0, 0
	aMap := make([]bool, len(guessStr)) // positions of a's
	countMap := make([]int, 10)         // occurences of chars (for calc b's)

	for i := range 10 {
		countMap[i] = strings.Count(player.Answer, strconv.Itoa(i))
	}

	log.Println(guessStr, player.Answer)
	for i := 0; i < len(guessStr); i++ {
		c, _ := strconv.Atoi(string(guessStr[i]))
		if guessStr[i] == player.Answer[i] {
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

	row := resultRow{
		TimeStamp: time.Now().Format(time.DateTime),
		Guess:     "#" + strconv.Itoa(len(player.GuessResults.Rows)+1) + ": " + guessStr,
		Result:    result,
	}

	player.GuessResults.Rows = append([]resultRow{row}, player.GuessResults.Rows...)

	h.renderer.Render(w, "result", player.GuessResults)
}
