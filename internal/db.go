package internal

import (
	"database/sql"
	"fmt"
)

type Record struct {
	Digits   int
	Name     string
	Attempts int
}

func InsertLeaderboard(data Record) error {
	db, err := sql.Open("sqlite3", "./leaderboards.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(fmt.Sprintf("insert into board%d(name, attempts) values(?, ?)", data.Digits))
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(data.Name, data.Attempts); err != nil {
		return err
	}

	return nil
}
