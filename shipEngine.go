package main

// This file contain all the meta to interract with or between ships

import (
	"math/rand"
)

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

// Checks if a position has a hit with existing boats
func hasCollission(boats []Ship, position ShipPosition) bool {
	for i := 0; i < len(boats); i++ {
		for j := 0; j < len(boats[i].positions); j++ {
			if position.x == boats[i].positions[j].x && position.y == boats[i].positions[j].y {
				return true
			}
		}
	}
	return false
}

// Generate a ship with the following pattern :
// Set the first position, then set a random orientation to generate remaining chained positions
func generateShip(ships []Ship, shipSize int8) Ship {
	var newShip Ship
	mustGenerate := true
	positions := make([]ShipPosition, 0, 5)
	var position ShipPosition
	const HORIZONTAL = 0
	const VERTICAL = 1

restartGeneration:
	// Is true when must generate/regenerate the entier ship
	for mustGenerate {

		// At the beginning, set the position slice empty
		positions = nil

		// Generate first position with x and y locations
		x := int8(rand.Intn(9 - int(shipSize)))
		y := int8(rand.Intn(9))

		// Regenerate the entire ship if has collision with other ships
		if hasCollission(ships, position) {
			break restartGeneration
		}

		// Append first position
		position.x = x
		position.y = y
		positions = append(positions, position)

		// Set a random orientation for the other parts of the ship
		orientation := int(rand.Intn(1))

		// Generate n remaining parts of the ship from the second index
		for i := 1; i < int(shipSize); i++ {
			if orientation == HORIZONTAL {
				position.x = x + int8(i)
				position.y = y
				positions = append(positions, position)

			} else if orientation == VERTICAL {
				position.x = x
				position.y = y + int8(i)
				positions = append(positions, position)
			}

			// Regenerate the entire ship if the position has a collision with existing position
			if hasCollission(ships, position) {
				break restartGeneration
			}
		}
	}

	// The positions are successfully generated
	newShip.positions = positions
	return newShip
}
