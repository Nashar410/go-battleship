package main

import (
	"fmt"
	"net/http"
)

// StartServer Start the server in go routines
func StartServer() {

}

// Return the board's state (case hit, ship touched, ship sinked)
func getBoard() {

}

// Return how many boat are left
func getBoats(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		remainingShip := len(ships)
		for _, ship := range ships {
			if shipSank(ship) {
				remainingShip--
			}
		}
		fmt.Fprintln(w, remainingShip)
	}
}

// Hit an opponent's ship's position
func postHit(idOpponent string, position ShipPosition) {
	if hasCollission(ships, position) {
		ships = hitShip(ships, position)
	}
}

func Send404NotFound(w http.ResponseWriter) {
	fmt.Fprintln(w, "Page not found")
	fmt.Println("Page not found")
}

func Send500InternalServerError(w http.ResponseWriter) {
	fmt.Fprintln(w, "Something went wrong")
	fmt.Println("Something went wrong")
}
