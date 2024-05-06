package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Game server started on port 8080:::")
	http.ListenAndServe(":8080")
}
