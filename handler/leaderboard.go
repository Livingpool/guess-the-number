package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/Livingpool/constants"
	"github.com/Livingpool/service"
	"github.com/Livingpool/views"
)

type LeaderboardHandler struct {
	renderer    views.TemplatesInterface
	leaderboard service.LeaderboardInterface
}

func NewLeaderboardHandler(r views.TemplatesInterface, l service.LeaderboardInterface) *LeaderboardHandler {
	return &LeaderboardHandler{
		renderer:    r,
		leaderboard: l,
	}
}

func (h *LeaderboardHandler) SaveRecord(w http.ResponseWriter, r *http.Request) {
	var record service.Record

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(&record); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Debug(fmt.Sprintf("decode json failed: %s", err.Error()))
		return
	}

	if record.Digits < constants.DIGIT_LOWER_LIMIT || record.Digits > constants.DIGIT_UPPER_LIMIT {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%d is out of range", record.Digits)))
		return
	}

	record.Name = strings.TrimSpace(record.Name)
	if len(record.Name) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("name too short"))
		return
	}

	if record.Attempts < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("attempts cannot < 1"))
		return
	}

	if err := h.leaderboard.Insert(r.Context(), record); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		slog.Error(fmt.Sprintf("insert leaderboard failed: %#v", err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("record inserted"))
}

func (h *LeaderboardHandler) ShowLeaderboard(w http.ResponseWriter, r *http.Request) {
	digit := r.URL.Query().Get("digit")
	name := strings.TrimSpace(r.URL.Query().Get("name"))

	boardId, err := strconv.Atoi(digit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%s is not a integer", digit)))
		return
	}

	if boardId < constants.DIGIT_LOWER_LIMIT || boardId > constants.DIGIT_UPPER_LIMIT {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("%d is out of range", boardId)))
		return
	}

	result, err := h.leaderboard.Get(r.Context(), boardId, name)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		slog.Debug(fmt.Sprintf("get leaderboard failed with id %d, %s: %v", boardId, name, err))
		w.Write([]byte("record not found"))
		return
	}

	h.renderer.Render(w, "leaderboard", result)
}
