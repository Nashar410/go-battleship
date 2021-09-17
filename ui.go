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

func generateEmptyBoard() [10][10]int8 {
	var board [10][10]int8
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			board[i][j] = 0
		}
	}
	return board
}

func DisplayOwnBoard(board []Ship) {
	fmt.Println("Generate a new board...")
	emptyBoard := generateEmptyBoard()
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i].positions); j++ {
			emptyBoard[board[i].positions[j].x][board[i].positions[j].y] = 1
		}
	}

	for x := 0; x < 10; x++ {
		fmt.Println(emptyBoard[x])
	}
	fmt.Println("Board successfully generated")
}

func displayOpponentBoard() {}

func displayOwnBoardState() {}
