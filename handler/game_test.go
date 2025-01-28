package handler_test

import (
	"context"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/Livingpool/handler"
	"github.com/Livingpool/middleware"
	"github.com/Livingpool/mocks"
	"github.com/Livingpool/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewGame(t *testing.T) {
	mockTemplates := mocks.NewMockTemplatesInterface(t)
	mockPlayerPool := mocks.NewMockPlayerPoolInterface(t)
	mockTimeProvider := mocks.NewMockTimeProviderInterface(t)
	mockGameHandler := handler.NewGameHandler(mockTemplates, mockPlayerPool, mockTimeProvider)

	testcases := []struct {
		name       string
		digit      string
		statusCode int
	}{
		{"alphabet_input", "a", 422},
		{"non_alnum_input", "!", 422},
		{"lower_than_range_input", "0", 422},
		{"higher_than_range_input", "15", 422},
		{"valid_input", "8", 200},
	}

	mockTemplates.EXPECT().Render(
		mock.AnythingOfType("*httptest.ResponseRecorder"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("service.FormData"),
	).RunAndReturn(func(w io.Writer, name string, data interface{}) error {
		funcMap := template.FuncMap{
			"dec": func(i int) int { return i - 1 },
		}
		newTemplates := template.Must(template.New("test").Funcs(funcMap).ParseGlob("./../views/html/*.tmpl"))
		newTemplates.ExecuteTemplate(w, name, data)
		return nil
	})

	mockPlayerPool.EXPECT().NewPlayer(mock.AnythingOfType("string")).Return(new(service.Player))
	mockPlayerPool.EXPECT().AddPlayer(mock.Anything).Return(nil)
	// mockPlayerPool.EXPECT().GetPlayer(mock.AnythingOfType("int")).Return(new(service.Player), true)
	// mockPlayerPool.EXPECT().RemovePlayer(mock.AnythingOfType("int")).Return(nil)

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/home?digit="+tc.digit, nil)

			r = r.WithContext(context.WithValue(r.Context(), middleware.RequestIdKey, uuid.New().String()))

			mockGameHandler.NewGame(w, r)

			assert.Equal(t, tc.statusCode, w.Code)
		})
	}
}

func TestCheckGuess(t *testing.T) {
	mockTemplates := mocks.NewMockTemplatesInterface(t)
	mockPlayerPool := mocks.NewMockPlayerPoolInterface(t)
	mockTimeProvider := mocks.NewMockTimeProviderInterface(t)
	mockGameHandler := handler.NewGameHandler(mockTemplates, mockPlayerPool, mockTimeProvider)

	testcases := []struct {
		name       string
		guess      string
		answer     string
		playerId   string
		result     string
		statusCode int
	}{
		{"empty_input", "", "123", "0", "", 422},
		{"wrong_length_input", "12", "123", "0", "", 422},
		{"wrong_id_format", "", "123", "a", "", 404},
		{"correct_guess", "111", "111", "0", "3a0b", 200},
		{"all_b", "1919", "9191", "0", "0a4b", 200},
		{"mixed_a_b", "1022", "1201", "0", "1a2b", 200},
	}

	mockTemplates.EXPECT().Render(
		mock.AnythingOfType("*httptest.ResponseRecorder"),
		mock.AnythingOfType("string"),
		mock.AnythingOfType("service.GuessResults"),
	).RunAndReturn(func(w io.Writer, name string, data interface{}) error {
		funcMap := template.FuncMap{
			"dec": func(i int) int { return i - 1 },
		}
		newTemplates := template.Must(template.New("test").Funcs(funcMap).ParseGlob("./../views/html/*.tmpl"))
		newTemplates.ExecuteTemplate(w, name, data)
		return nil
	})

	// i dunno why this failed...
	// mockTemplates.On(
	// 	"Render",
	// 	mock.AnythingOfType("*httptest.ResponseRecorder"),
	// 	mock.AnythingOfType("string"),
	// 	mock.AnythingOfType("service.GuessResults"),
	// ).Return(func(w io.Writer, name string, data interface{}) error {
	// 	funcMap := template.FuncMap{
	// 		"dec": func(i int) int { return i - 1 },
	// 	}
	// 	newTemplates := template.Must(template.New("example").Funcs(funcMap).ParseGlob("./../views/html/*.tmpl"))
	// 	newTemplates.ExecuteTemplate(w, name, data)
	// 	return nil
	// }).Times(len(testcases))

	mockTimeProvider.EXPECT().Now(mock.Anything).Return(time.Now())

	for _, tc := range testcases {
		if tc.statusCode != 404 {
			newPlayer := new(service.Player)
			newPlayer.Answer = tc.answer
			newPlayer.Id, _ = strconv.Atoi(tc.playerId)
			mockPlayerPool.On(
				"GetPlayer",
				newPlayer.Id,
			).Return(newPlayer, true).Once()
		}
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/check?guess="+tc.guess+"&id="+tc.playerId, nil)

			r = r.WithContext(context.WithValue(r.Context(), middleware.RequestIdKey, uuid.New().String()))

			mockGameHandler.CheckGuess(w, r)

			assert.Equal(t, tc.statusCode, w.Code)
		})
	}
}
