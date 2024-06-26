# Guess The Number
This is my first project written in Go & HTMX and I quite enjoyed it. \
It's a simple game where a hint is given to each guess until the user gets the correct answer.

<p align="center">
<img height="400" width="320" alt="Screenshot 2024-05-31 at 5 14 29 PM" src="https://github.com/Livingpool/guess-the-number/assets/52132459/97dba180-16b1-4f25-b29c-48489c7b872f">
&nbsp;&nbsp;&nbsp;
<img height="400" width="350" alt="Screenshot 2024-05-31 at 5 13 26 PM" src="https://github.com/Livingpool/guess-the-number/assets/52132459/704bee09-fcac-4242-8765-db64e18e7a59">
</p>

### To run this project in development
 - Old school way (Go version 1.22 or above is required):
```bash
cd cmd
go run .
```
- If you have [Air](https://github.com/cosmtrek/air) installed
```bash
 air -c ./air.toml
```

### Project structure
When a client hits the /new endpoint, a new player instance is created and added to the player pool. \
A player will automatically be removed from the pool after some time of inactivity(timeout).
```
  ├── cmd
        ├── main.go           
  ├── constants                  (constants, auto-incremented playerId)
  ├── handler
        ├── testdata             (test data directory auto-generated by goldie, for testing templates)
        ├── game.go              (main game logic and http handlers)
        ├── game_test.go         (unit tests using testify, mockery, and goldie)
        ├── player.go            (definitions for Player)
        ├── playerPool.go        (definitions for PlayerPool)
  ├── middleware
  ├── mocks                      (mocks auto-generated by mockery)
  ├── router                     (routing based on Go 1.22 routing enhancements)
  ├── views
        ├── assets
        ├── css
        ├── html
        ├── scripts              (additional JS code for better visuals)
        ├── html.go              (code for rendering templates)           
```

