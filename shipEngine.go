package main

// Position of a ship
type ShipPosition struct {
	x int8
	y int8
}

// A ship
type Ship struct {
	// Positions for this ship
	positions []ShipPosition
	// Positions where the ship is touched
	touchedAt []ShipPosition
}
