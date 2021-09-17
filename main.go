package main

import "fmt"

func main() {
	ships := generateShips()
	DisplayOwnBoard(ships)
	fmt.Println()
}

func generateShips() []Ship {
	ships := make([]Ship, 0, 2)

	// Add Two boats
	ships = append(ships, generateShip(ships, 2))
	ships = append(ships, generateShip(ships, 3))
	return ships
}
