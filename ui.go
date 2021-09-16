package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//////////////////////////////////////////////////
////////// 			GENERIC 			//////////
//////////////////////////////////////////////////
const BOARD_SIZE = 5
const COMMAND_NO string = "n"
const COMMAND_YES string = "y"
const COMMAND_QUIT string = "q"
const COMMAND_ATTACK string = "a"
const COMMAND_SEE_OWN_BOARD string = "z"
const COMMAND_SEE_OPPONENT_BOARD string = "o"
const COMMAND_SEE_OWN_BOARD_STATE string = "s"

func askPlayer(question string) string {
	fmt.Println(question)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erreur, réessayez !", err)
		return ""
	}
	input = strings.TrimSuffix(input, "\n")
	return input
}

// Create an array of int8 array that will be fill
func generateEmptyBoard() [BOARD_SIZE][]int8 {
	var board [BOARD_SIZE][]int8
	for i := 0; i < BOARD_SIZE; i++ {
			board[i] = make([]int8, BOARD_SIZE)
	}
	return board
}

var Board []Ship = []Ship{
	{
		[]ShipPosition{{0, 0}},
		[]ShipPosition{{}}},
	{
		[]ShipPosition{{0, 2}},
		[]ShipPosition{{}}},
	{
		[]ShipPosition{{1, 1}},
		[]ShipPosition{{}}}}

func GenerateAndShowABoard(board []Ship) {
	showABoard(fillABoard(board))
}

func fillABoard(board []Ship) [BOARD_SIZE][]int8 {
	generatedBoard := generateEmptyBoard()
	for _, ship := range board {
		for _, shipPosition := range ship.Positions {
			generatedBoard[shipPosition.x][shipPosition.y] = 1
		}
	}
	return generatedBoard
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
