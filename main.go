package main

import (
	"os"
)

var idPlayer = RandomString(4)
var idsOpponent []string
var caseTouched []ShipPosition
var ships []Ship
var ownPort string = "8080"

func main() {
	ownPort = os.Args[1]
	StartServer()
	ships := generateShips()
	GenerateAndShowABoard(ships)
	WelcomePlayer(idPlayer)
	ActionMenu()
}

func generateShips() []Ship {
	// Add Two boats
	ships = append(ships, generateShip(ships, 2))
	ships = append(ships, generateShip(ships, 3))
	ships = append(ships, generateShip(ships, 4))
	return ships
}
