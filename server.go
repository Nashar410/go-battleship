package main

import (
	"fmt"
	"net/http"
)

// StartServer Start the server in go routines
func StartServer() {
	http.HandleFunc("/board", getBoard)
	http.HandleFunc("/boats", getBoats)
	http.HandleFunc("/hit", postHit)
	http.ListenAndServe(":8080", nil)
}

// Return the board's state (case hit, ship touched, ship sinked)
func getBoard(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		boards := fillAStateBoard(ships, caseTouched)
		fmt.Printf("State of the board : \n")
		fmt.Fprintf(w, "State of the board : \n")
		for _, board := range boards {
			fmt.Printf("%d \n", board)
			fmt.Fprintf(w, "%d \n", board)
		}

	}
}

// Return how many boat are remaining
func getBoats(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		fmt.Fprintln(w, getSurvivingShips(ships))
	default:
		Send404NotFound(w)
	}
}

// Hit an opponent's ship's position
func postHit(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			Send500InternalServerError(w)
			return
		}
		// We check the form keys
		if !req.Form.Has("x") || !req.Form.Has("y") {
			send500InvalidForm(w)
			return
		}

		// Make a position from the user data
		var x, y int8
		fmt.Sscanf(req.Form.Get("x"), "%d", x)
		fmt.Sscanf(req.Form.Get("y"), "%d", y)
		var position = ShipPosition{x: x, y: y}

		if hasCollission(ships, position) {
			ships = hitShip(ships, position)
		} else {
			caseTouched = append(caseTouched, position)
		}

		// Send the success state to the user
		fmt.Fprintf(w, "%s\n", "HIT SENDED")
	default:
		Send404NotFound(w)
	}
}

func Send404NotFound(w http.ResponseWriter) {
	fmt.Fprintln(w, "Page not found")
}

func Send500InternalServerError(w http.ResponseWriter) {
	fmt.Fprintln(w, "Something went wrong")
}

func send500InvalidForm(w http.ResponseWriter) {
	fmt.Fprintln(w, "Form invalid")
}
