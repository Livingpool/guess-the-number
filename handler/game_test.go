package handler_test

import (
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/Livingpool/handler"
	"github.com/Livingpool/mocks"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewGame(t *testing.T) {
	mockTemplates := &mocks.MockTemplatesInterface{}
	mockPlayerPool := &mocks.MockPlayerPoolInterface{}
	mockTimeProvider := &mocks.MockTimeProviderInterface{}
	mockGameHandler := handler.NewGameHandler(mockTemplates, mockPlayerPool, mockTimeProvider)

	testcases := []struct {
		name       string
		digit      string
		statusCode int
	}{
		{"alphabet input", "a", 422},
		{"non alnum input", "!", 422},
		{"lower than range input", "0", 422},
		{"higher than range input", "15", 422},
		{"valid input", "8", 200},
	}

	mockTemplates.On(
		"Render",
		mock.AnythingOfType("*httptest.ResponseRecorder"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("handler.FormData"),
	).Return(func(w io.Writer, name string, data interface{}) error {
		newTemplates := template.Must(template.ParseGlob("./../views/html/*.html"))
		newTemplates.ExecuteTemplate(w, name, data)
		return nil
	}).Times(len(testcases))

	mockPlayerPool.On(
		"NewPlayer",
		mock.AnythingOfType("string"),
	).Return(new(handler.Player)).Times(len(testcases))

	mockPlayerPool.On(
		"AddPlayer",
		mock.Anything,
	).Return(nil).Times(len(testcases))

	g := goldie.New(t)

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/home?digit="+tc.digit, nil)

			mockGameHandler.NewGame(w, r)

			assert.Equal(t, tc.statusCode, w.Code)
			g.Assert(t, tc.name, w.Body.Bytes())
		})
	}
}

func TestCheckGuess(t *testing.T) {
	mockTemplates := &mocks.MockTemplatesInterface{}
	mockPlayerPool := &mocks.MockPlayerPoolInterface{}
	mockTimeProvider := &mocks.MockTimeProviderInterface{}
	mockGameHandler := handler.NewGameHandler(mockTemplates, mockPlayerPool, mockTimeProvider)

	testcases := []struct {
		name       string
		guess      string
		answer     string
		playerId   string
		result     string
		statusCode int
	}{
		{"empty input", "", "123", "0", "", 422},
		{"wrong length input", "12", "123", "0", "", 422},
		{"wrong id format", "", "123", "a", "", 404},
		{"correct guess", "111", "111", "0", "3a0b", 200},
		{"all b", "1919", "9191", "0", "0a4b", 200},
		{"mixed a b", "1022", "1201", "0", "1a2b", 200},
	}

	mockTemplates.On(
		"Render",
		mock.AnythingOfType("*httptest.ResponseRecorder"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("handler.guessResults"),
	).Return(func(w io.Writer, name string, data interface{}) error {
		newTemplates := template.Must(template.ParseGlob("./../views/html/*.html"))
		newTemplates.ExecuteTemplate(w, name, data)
		return nil
	}).Times(len(testcases))

	mockTimeProvider.On(
		"Now",
	).Return(time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC))

	for _, tc := range testcases {
		if tc.statusCode != 404 {
			newPlayer := new(handler.Player)
			newPlayer.Answer = tc.answer
			newPlayer.Id, _ = strconv.Atoi(tc.playerId)
			mockPlayerPool.On(
				"GetPlayer",
				newPlayer.Id,
			).Return(newPlayer, true).Once()
		}
	}

	g := goldie.New(t)

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/check?guess="+tc.guess+"&id="+tc.playerId, nil)

			mockGameHandler.CheckGuess(w, r)

			timeData := struct {
				Time string
			}{
				time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).Format(time.DateTime),
			}

			assert.Equal(t, tc.statusCode, w.Code)
			g.AssertWithTemplate(t, tc.name, timeData, w.Body.Bytes())
		})
	}
}
