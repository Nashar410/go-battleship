package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var WELCOME_PLAYER = Choice{
	"Bienvenue au jeu de bataille navale !\n",
	[]string{}}

var GIVE_PLAYER_ID = Choice{
	"Votre ID est %s\n",
	[]string{}}

var ENTER_PORT_OPPONENT = Choice{
	"Entrez l'ID d'un adversaire:\n ",
	[]string{}}

var ENTER_MORE_PORT_OPPONENT = Choice{
	"Voulez-vous ajouter un autre adversaire ?\n",
	[]string{
		COMMAND_YES + " pour oui\n",
		COMMAND_NO + " pour non\n"}}

var RULES = Choice{
	"",
	[]string{}}

var PERSONNAL_ACTION_MENU = Choice{
	"Voici les actions possibles:\n ",
	[]string{
		COMMAND_SEE_OWN_BOARD + " pour voir la position de vos navires\n",
		COMMAND_SEE_OWN_BOARD_STATE + " pour voir l'état de vos navires\n"}}

var ACCESS_OPPONENT_ACTION_MENU = Choice{
	"Accéder au menu de combat\n",
	[]string{
		COMMAND_ACTION_OPPONENT_BUTTON + " pour attaquer un adversaire\n"}}

var OWN_BOAT_POSITION = Choice{
	"Etats des navires\n",
	[]string{}}

var OWN_BOAT_STATE = Choice{
	"Positions des navires\n",
	[]string{}}

var COMBAT_MENU = Choice{
	"Menu de combat\n",
	[]string{}}

var LEGEND_STATE_BOARD = Choice{
	"Légende:\n",
	[]string{
		"0 ⇥ Case non-touchée\n",
		"1 ⇥ Navire touché sur cette case\n",
		"2 ⇥ Navire coulé\n",
		"9 ⇥ Case touchée\n"}}

var LEGEND_POSITION_BOARD = Choice{
	"★ Légende:\n",
	[]string{
		"0 ⇥ Case vide\n",
		"1 ⇥ 1 Navire\n",
		"Plusieurs 1 adjacents ⇥ Navire de plusieurs cases\n",
	}}
var ENTER_CHOICE = Choice{
	"★ Entrez un choix:\n",
	[]string{}}

var PRESS_TO_CONTINUE = Choice{
	"Appuyez sur entrée pour continuer\n",
	[]string{}}

var BOAT_STATE = Choice{
	"Appuyez sur entrée pour continuer\n",
	[]string{}}

var UNEXPECTED_ACTION = Choice{
	"Action non prévue\n",
	[]string{}}

var WHICH_OPPONENT = Choice{
	"Choisissez votre cible:\n ",
	[]string{}}

var OPPONENT_ACTION_MENU = Choice{
	"Voici les actions possibles:\n ",
	[]string{
		COMMAND_ATTACK + " pour attaquer %s\n",
		COMMAND_SEE_OPPONENT_BOARD + " pour voir l'état du plateau de %\n"}}

var WHICH_OPPONENT_CASE = Choice{
	"Quel case voulez-vous attaquer ? Format attendu: 1:1\n",
	[]string{}}

var ATTACK_LAUNCHED = Choice{
	"Vous avez lancé une attaque sur %s\n",
	[]string{}}

var ATTACK_SUCCESSED = Choice{
	"Vous avez touché un des navires de %s\n",
	[]string{}}

var ATTACK_FAILED = Choice{
	"Vous n'avez touché aucun des navires de %s\n",
	[]string{}}

var ATTACKED_SUCCESSED = Choice{
	"Un de vos navires a été touché par un tir de %s\n",
	[]string{}}

var ATTACKED_FAILED = Choice{
	"Vous avez essuyé un tir de %s sans dégats %s\n",
	[]string{}}

var ATTACKED_SINKED = Choice{
	"Vous n'avez touché aucun des navires de %s, un de vos navires a coulé\n",
	[]string{}}

var YOU_LOST = Choice{
	"Tout vos navires ont été coulés, vous avez perdu.\n",
	[]string{}}

var OPPONENT_LOST = Choice{
	"%s n'a plus aucun navire, il a perdu.\n",
	[]string{}}

var YOU_WIN = Choice{
	"Vous êtes le dernier en lice ! Vous avez gagné !\n",
	[]string{}}

type Choice struct {
	Text    string
	Choices []string
}

func WelcomePlayer() {
	// Welcome the player
	fmt.Println(WELCOME_PLAYER.getText())
	// Ask for an opponent id
	reAskportOpponent := true
	firstAsk := true
	for reAskportOpponent {
		if firstAsk {
			portsOpponent = append(portsOpponent, askPlayer(ENTER_PORT_OPPONENT.getText()))
			firstAsk = false
		} else {
			reAskportOpponent = askPlayer(ENTER_MORE_PORT_OPPONENT.getTextWithChoices()) == COMMAND_YES
			if reAskportOpponent {
				portsOpponent = append(portsOpponent, askPlayer(ENTER_PORT_OPPONENT.getText()))
			}
		}

	}
	fmt.Println(RULES.getText() + "rules")
}

func ActionMenu() {
	for gameContinue {
		fmt.Println(PERSONNAL_ACTION_MENU.getTextWithChoices())
		fmt.Println(ACCESS_OPPONENT_ACTION_MENU.getTextWithChoices())
		playerChoice := askPlayer(ENTER_CHOICE.getText())
		switch playerChoice {
		case COMMAND_SEE_OWN_BOARD:
			fmt.Println(OWN_BOAT_POSITION.getText())
			showOwnBoard()
			fmt.Println(LEGEND_POSITION_BOARD.getTextWithChoices())
		case COMMAND_SEE_OWN_BOARD_STATE:
			fmt.Println(OWN_BOAT_STATE.getText())
			showOwnBoardState()
			fmt.Println(LEGEND_STATE_BOARD.getTextWithChoices())

		case COMMAND_ACTION_OPPONENT_BUTTON:
			OpponentActionMenu()
		default:
			fmt.Println(UNEXPECTED_ACTION.getText())
		}

		if areAllSink(ships) {
			gameContinue = false
			fmt.Println(YOU_LOST.getText())
		}

		askPlayer(PRESS_TO_CONTINUE.getText())
	}
}

func OpponentActionMenu() {
	fmt.Println(COMBAT_MENU.getText())
	portOpponent := askPlayer(WHICH_OPPONENT.getTextWithChoices())
	fmt.Println(OPPONENT_ACTION_MENU.getTextWithChoices())
	choicePlayer := askPlayer(ENTER_CHOICE.getText())
	switch choicePlayer {
	case COMMAND_ATTACK:
		fmt.Println(portOpponent)
	case COMMAND_SEE_OPPONENT_BOARD:
		fmt.Println(portOpponent)
	default:
		fmt.Println(UNEXPECTED_ACTION.getText())
	}
	ActionMenu()
}

// rajoute un joueur
func (choiceObj Choice) getTextWithChoices() string {
	choices := ""
	for _, choice := range choiceObj.Choices {
		choices += "\t  ▶ " + choice
	}
	return choiceObj.getText() + choices
}

// rajoute un joueur
func (choiceObj Choice) getText() string {
	return "⚡\t" + choiceObj.Text
}

func askPlayer(question string) string {
	fmt.Println(question)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erreur, réessayez !", err)
		return ""
	}
	if strings.Contains(input, "\r") {
		return strings.TrimSuffix(input, "\r\n")
	} else {
		return strings.TrimSuffix(input, "\n")
	}
}

/*
Scénario :

Bienvenue au jeu de battle
Votre Id est : %s (string{A.a.9})
Entrez l'id d'un autre joueur :
Voulez-vous entrer l'id d'un autre joueur ?
...Règle du jeu
Action possible :
- Voir son board ou son state de board
- Sautez une ligne
- Action sur Id joueur :
	- Voir son state (tableau ou alors boolean false = joueur ko)
	- Attaquer ce joueur
[Id Joueur] a perdu
Vous avez reçu un essuyé un tir qui a touché [bateau] (afficher state board)
Vous avez reçu un essuyé un tir qui a coulé [bateau] (afficher state board)
Vous avez reçu un essuyé un tir qui n'a rien coulé
Tout vos bateaux ont été coulé, vous avez perdu
Votre (vos) adversaire(s) n'a (ont) plus de bateau à flot, vous êtes le vainqueur !
*/
