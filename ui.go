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

func displayToUser(str string) {
	fmt.Println(str)
}

func generateEmptyBoard() [][]int8 {
	var board [][]int8
	for i := 0; i < 10; i++ {
		for y := 0; y < 10; y++ {
			board[i][y] = 0
			// utiliser make !
		}
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

func DisplayOwnBoard(board []Ship) {
	emptyBoard := generateEmptyBoard()
	for i := 0; i < len(board); i++ {
		for y := 0; y < len(board[i].positions); y++ {
			emptyBoard[board[i].positions[y].x][board[i].positions[y].y] = 1
		}
		fmt.Println(emptyBoard[i])
	}
}

func displayOpponentBoard() {}

func displayOwnBoardState() {}
