package constants

import "sync"

const (
	PLAYER_POOL_CAP int = 3
)

var (
	AutoInc autoInc
)

type autoInc struct {
	sync.Mutex
	id int
}

// Function for generating auto incrementing id.
func (a *autoInc) ID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id
	a.id++
	return
}
