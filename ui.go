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
	if input == COMMAND_QUIT {
		return ""
	} else {
		return input
	}
}
