package main

import (
	"fmt"
)

//////////////////////////////////////////////////
////////// 			GENERIC 			//////////
//////////////////////////////////////////////////
const BOARD_SIZE = 10
const COMMAND_NO = "n"
const COMMAND_YES = "y"
const COMMAND_QUIT = "q"
const COMMAND_ATTACK = "a"
const COMMAND_ACTION_OPPONENT_BUTTON = "w"
const COMMAND_SEE_OWN_BOARD = "z"
const COMMAND_SEE_OPPONENT_BOARD = "o"
const COMMAND_SEE_OWN_BOARD_STATE = "s"
const CASE_TOUCHED = 9
const SHIP_TOUCHED = 1
const SHIP_SINKED = 2
const CASE_EMPTY = 0

// Create an array of int8 array that will be fill
func generateEmptyBoard() [BOARD_SIZE][]int8 {
	var board [BOARD_SIZE][]int8
	for i := 0; i < BOARD_SIZE; i++ {
		board[i] = make([]int8, BOARD_SIZE)
	}
	return board
}

func showOwnBoard() {
	GenerateAndShowABoard(ships)
}

func showOwnBoardState() {
	showABoard(fillAStateBoard(ships, caseTouched))
}

func GenerateAndShowABoard(board []Ship) {
	showABoard(fillABoard(board))
}

func fillABoard(board []Ship) [BOARD_SIZE][]int8 {
	generatedBoard := generateEmptyBoard()
	for _, ship := range board {
		for _, shipPosition := range ship.positions {
			generatedBoard[shipPosition.x][shipPosition.y] = 1
		}
	}
	return generatedBoard
}

// Generate a board where the state of the ships and case are rendered
func fillAStateBoard(board []Ship, caseTouched []ShipPosition) [BOARD_SIZE][]int8 {
	generatedBoard := generateEmptyBoard()

	// If a case without ship already touched
	if len(caseTouched) > 0 {
		for _, oneCase := range caseTouched {
			generatedBoard[oneCase.x][oneCase.y] = CASE_TOUCHED

		}
	}

	for _, ship := range board {
		// If ship sinked
		if shipSank(ship) {
			for _, shipPosition := range ship.positions {
				generatedBoard[shipPosition.x][shipPosition.y] = SHIP_SINKED
			}
		} else if shipWasHit(ship) {
			for _, shipTouchedAt := range ship.touchedAt {
				generatedBoard[shipTouchedAt.x][shipTouchedAt.y] = SHIP_TOUCHED
			}
		}

	}

	return generatedBoard
}

func shipSank(ship Ship) bool {
	return len(ship.positions) == len(ship.touchedAt)
}

func shipWasHit(ship Ship) bool {
	return len(ship.touchedAt) > 0
}

func showABoard(board [BOARD_SIZE][]int8) {
	for _, line := range board {
		fmt.Println(line)
	}
}

/*

// Display the board of the player
func DisplayOwnBoard(board []Ship) {
}
// Display opponent board
func DisplayOpponentBoard() {}

// Display own state board
func DisplayOwnBoardState() {}


*/
