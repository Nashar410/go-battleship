package main

import "fmt"

var idPlayer = RandomString(4)
var idsOpponent []string
var caseTouched []ShipPosition

func main() {
	ships := generateShips()
	GenerateAndShowABoard(ships)
	fmt.Println()
}

func generateShips() []Ship {
	ships := make([]Ship, 0, 3)

	// Add Two boats
	ships = append(ships, generateShip(ships, 2))
	ships = append(ships, generateShip(ships, 3))
	ships = append(ships, generateShip(ships, 4))
	return ships
}
