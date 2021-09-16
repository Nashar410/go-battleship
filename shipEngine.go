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

// Min x 0
// Min y 0
// Max x 9
// Max y 9

func createGame() {
	// ships := make([]Ship, 0, 5)

	// ship = generateShip(ships, 2)
	// ship = generateShip(ships, 3)
}

func generateShip(ships []Ship, shipSize int8) Ship {
	var ship Ship
	positions := make([]ShipPosition, 0, 5)
	for i := 0; i < int(shipSize); i++ {
		var position ShipPosition
		searchPosition := true
		for searchPosition {
			// Generate coord

			// Verify collission
			if !hasCollission(ships, position.x, position.y) {
				searchPosition = false
			}
		}
		positions = append(positions, position)
	}

	ship.positions = positions
	return ship
}
