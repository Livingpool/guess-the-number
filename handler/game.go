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
	renderer     views.TemplatesInterface
	playerPool   PlayerPoolInterface
	timeProvider TimeProviderInterface
}

func NewGameHandler(renderer views.TemplatesInterface, playerPool PlayerPoolInterface, timeProvider TimeProviderInterface) *GameHandler {
	return &GameHandler{
		renderer:     renderer,
		playerPool:   playerPool,
		timeProvider: timeProvider,
	}
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
		w.Header().Set("HX-Retarget", "#form")
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.renderer.Render(w, "form", formData)
		return
	}

	// Input value not in range error
	if digit < 1 || digit > 8 {
		formData := FormData{
			Error: "Input is not in range :(",
		}
		w.Header().Set("HX-Retarget", "#form")
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.renderer.Render(w, "form", formData)
		return
	}

	lower, upper := calcRange(digit)
	answer := genRandInt(lower, upper)
	answerStr := strconv.Itoa(answer)

	// PlayerPool full error
	newPlayer := h.playerPool.NewPlayer(answerStr)
	if err = h.playerPool.AddPlayer(newPlayer); err != nil {
		formData := FormData{
			Error: "Server is full. Please try again later!",
		}
		w.Header().Set("HX-Retarget", "#form")
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.renderer.Render(w, "form", formData)
		return
	}

	formData := FormData{
		Digit:    digit,
		Start:    strconv.Itoa(lower),
		End:      strconv.Itoa(upper),
		Error:    "",
		PlayerId: newPlayer.Id,
	}

	log.Println("Player registered", newPlayer.Id)

	// Execute the template
	h.renderer.Render(w, "game", formData)
}

func (h *GameHandler) CheckGuess(w http.ResponseWriter, r *http.Request) {
	guessStr := r.URL.Query().Get("guess")
	playerId := r.URL.Query().Get("id")

	// Player id not parseable error
	id, err := strconv.Atoi(playerId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Player doesn't exist error
	player, exists := h.playerPool.GetPlayer(id)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Guess/Answer length not match error
	if len(guessStr) != len(player.Answer) {
		w.WriteHeader(http.StatusUnprocessableEntity)
		h.renderer.Render(w, "result", player.GuessResults)
		return
	}

	a, b := 0, 0
	aMap := make([]bool, len(guessStr)) // positions of a's
	countMap := make([]int, 10)         // occurences of chars (for calc b's)

	for i := range 10 {
		countMap[i] = strings.Count(player.Answer, strconv.Itoa(i))
	}

	// log.Println(guessStr, player.Answer)
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
		TimeStamp: h.timeProvider.Now().Format(time.DateTime),
		Guess:     "#" + strconv.Itoa(len(player.GuessResults.Rows)+1) + ": " + guessStr,
		Result:    result,
	}

	player.GuessResults.Rows = append([]resultRow{row}, player.GuessResults.Rows...)

	h.renderer.Render(w, "result", player.GuessResults)
}

// Returns [lower, upper] from given digit
func calcRange(digit int) (int, int) {
	if digit == 1 {
		return 0, 9
	}
	upper, _ := strconv.Atoi(strings.Repeat("9", digit))
	return int(math.Pow(10, float64(digit-1))), upper
}

// Generate a pseudo random number with range [lower, upper]
func genRandInt(lower, upper int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(upper-lower) + lower
	return num
}
