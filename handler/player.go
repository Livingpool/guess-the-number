package handler

import "time"

type TimeProviderInterface interface {
	Now() time.Time
}

type RealTimeProvider struct{}

func (rtp *RealTimeProvider) Now() time.Time {
	return time.Now()
}

type Player struct {
	Id           int
	Answer       string
	GuessResults guessResults
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
