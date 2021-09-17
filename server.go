package main

import (
	"fmt"
	"net/http"
)

// StartServer Start the server in go routines
func StartServer() {
	http.HandleFunc("/board", getBoard)
	http.HandleFunc("/boats", getBoats)
	//http.HandleFunc("/hits", postHit)
	http.ListenAndServe(":8080", nil)

}

// Return the board's state (case hit, ship touched, ship sinked)
func getBoard(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Printf("%s", fillABoard(ships))
	}
}

// Return how many boat are remaining
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
	default:
		Send404NotFound(w)
	}
}

// Hit an opponent's ship's position
//func postHit(w http.ResponseWriter, req *http.Request) {
//	//if hasCollission(ships, position) {
//	//	ships = hitShip(ships, position)
//	//}
//}

func Send404NotFound(w http.ResponseWriter) {
	fmt.Fprintln(w, "Page not found")
}

func Send500InternalServerError(w http.ResponseWriter) {
	fmt.Fprintln(w, "Something went wrong")
}
