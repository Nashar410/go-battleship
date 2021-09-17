package main

// This file contain all the meta to interract with or between ships

import (
	"fmt"
	"math/rand"
	"time"
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

func hitShip(ships []Ship, position ShipPosition) []Ship {
	for _, ship := range ships {
		for _, curPosition := range ship.positions {
			if curPosition == position {
				ship.touchedAt = append(ship.touchedAt, position)
				return ships
			}
		}
	}
	return ships
}

// Generate a ship with the following pattern :
// Set the first position, then set a random orientation to generate remaining chained positions
func generateShip(ships []Ship, shipSize int8) (newShip Ship) {
	mustGenerate := true
	positions := make([]ShipPosition, 0, shipSize)
	var position ShipPosition
	const HORIZONTAL = 0
	const VERTICAL = 1

	fmt.Println("Generate a new ship...")
	for mustGenerate {
		// At the beginning, set the position slice empty
		positions = nil

		// Set a random orientation for the other parts of the ship
		rand.Seed(time.Now().Unix())
		orientation := int(rand.Intn(10)) // 5/10 chance => "1/2 chance"

		// Generate first position with x and y locations
		rand.Seed(time.Now().Unix())
		var x int8
		var y int8
		if orientation%2 == HORIZONTAL {
			x = int8(rand.Intn(9 - int(shipSize)))
			y = int8(rand.Intn(9))
		} else {
			x = int8(rand.Intn(9))
			y = int8(rand.Intn(9 - int(shipSize)))
		}
		position.x = x
		position.y = y

		// Regenerate the entire ship if has collision with other ships
		if !hasCollission(ships, position) {
			// Append first position
			positions = append(positions, position)

			// Generate n remaining parts of the ship from the second index
			i := int8(1)
			for i < shipSize {
				if orientation%2 == HORIZONTAL {
					position.x = x + int8(i)
					position.y = y
					positions = append(positions, position)

				} else if orientation%2 == VERTICAL {
					position.x = x
					position.y = y + int8(i)
					positions = append(positions, position)
				}

				// Regenerate the entire ship if the position has a collision with existing position
				if hasCollission(ships, position) {
					break
				}
				i++
			}
			if i == shipSize {
				mustGenerate = false
				break
			}

		}
		// Collision !
		time.Sleep(time.Second)
	}

	// The positions are successfully generated
	fmt.Println("Ship successfully generated")
	newShip.positions = positions
	return
}

func shipSank(ship Ship) bool {
	return len(ship.positions) == len(ship.touchedAt)
}

func shipHadHit(ship Ship) bool {
	return len(ship.touchedAt) > 0
}

func areAllSink(ships []Ship) bool {
	return getSurvivingShips(ships) == 0
}

func getSurvivingShips(ship []Ship) (remainingShip int) {
	remainingShip = len(ships)
	for _, ship := range ships {
		if shipSank(ship) {
			remainingShip--
		}
	}
	return
}
