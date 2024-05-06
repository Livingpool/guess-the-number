package logic

import (
	"math"
	"math/rand"
	"strconv"
	"strings"
)

func NewGame(digits int) string {
	answer := rand.Intn(int(math.Pow(10, float64(digits-1))))
	answerStr := strconv.Itoa(answer)
	for i := 0; i < digits-len(answerStr); i++ {
		answerStr = "0" + answerStr
	}

	return answerStr
}

func CheckGuess(guess, answer string) string {
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
