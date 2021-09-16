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

// Checks if a shot has a hit
func hasCollission(boats []Ship, x int8, y int8) bool {
	for i := 0; i < len(boats); i++ {
		for j := 0; j < len(boats[i].positions); j++ {
			if x == boats[i].positions[j].x && y == boats[i].positions[j].y {
				return true
			}
		}
	}
	return false
}
