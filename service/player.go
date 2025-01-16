package service

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
	GuessResults GuessResults
}

type FormData struct {
	Digit    int
	Start    string
	End      string
	Error    string
	PlayerId int
}

type ResultRow struct {
	TimeStamp string
	Guess     string
	Result    string
}

type GuessResults struct {
	Rows []ResultRow
}
