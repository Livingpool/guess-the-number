package handler

import (
	"errors"
	"sync"

	"github.com/Livingpool/constants"
)

type Player struct {
	Id           int
	GuessResults guessResults
	Answer       string
}

func NewPlayer(answer string) *Player {
	return &Player{
		Id:     constants.AutoInc.ID(),
		Answer: answer,
	}
}

type PlayerPool struct {
	players  map[int]*Player
	lock     *sync.RWMutex
	capacity int
}

func NewPlayerPool(capacity int) *PlayerPool {
	return &PlayerPool{
		players:  make(map[int]*Player),
		lock:     new(sync.RWMutex),
		capacity: capacity,
	}
}

func (p *PlayerPool) AddPlayer(player *Player) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	// If pool is at capacity, return error
	if len(p.players) == p.capacity {
		return errors.New("Player pool at capacity")
	}

	p.players[player.Id] = player
	return nil
}

func (p *PlayerPool) RemovePlayer(id int) error {
	p.lock.Lock()
	defer p.lock.Unlock()

	delete(p.players, id)
	return nil
}
