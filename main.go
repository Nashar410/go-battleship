package main

import "fmt"

func main() {
	fmt.Println()
}


// Position of a ship
type ShipPosition struct {
	x int8
	y int8
}

// A ship
type Ship struct {
	// Positions for this ship
	Positions []ShipPosition
	// Positions where the ship is touched
	TouchedAt []ShipPosition
}
