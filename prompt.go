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

var ENTER_ID_OPPONENT = Choice{
	"Entrez l'ID d'un adversaire: ",
	[]string{}}

var ENTER_MORE_ID_OPPONENT = Choice{
	"Voulez-vous ajouter un autre adversaire ?",
	[]string{
		COMMAND_YES + " pour oui\n",
		COMMAND_NO + " pour non\n"}}

var RULES = Choice{
	"",
	[]string{}}

var PERSONNAL_ACTION_MENU = Choice{
	"Voici les actions possibles: ",
	[]string{
		COMMAND_SEE_OWN_BOARD + " pour voir la position de vos navires\n",
		COMMAND_SEE_OWN_BOARD_STATE + " pour voir l'état de vos navires\n"}}

var ACCESS_OPPONENT_ACTION_MENU = Choice{
	"Accéder au menu de combat\n",
	[]string{
		COMMAND_ACTION_OPPONENT_BUTTON + " pour attaquer un adversaire\n"}}

var COMBAT_MENU = Choice{
	"Menu de combat\n",
	[]string{}}

var ENTER_CHOICE = Choice{
	"Entrez un choix:\n",
	[]string{}}

var UNEXPECTED_ACTION = Choice{
	"Action non prévue\n",
	[]string{}}

var WHICH_OPPONENT = Choice{
	"Choisissez votre cible: ",
	[]string{}}

var OPPONENT_ACTION_MENU = Choice{
	"Voici les actions possibles: ",
	[]string{
		COMMAND_ATTACK + " pour attaquer %s\n",
		COMMAND_SEE_OPPONENT_BOARD + " pour voir l'état du plateau de %\n"}}

var WHICH_OPPONENT_CASE = Choice{
	"Quel case voulez-vous attaquer ? Format attendu: 1:1",
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

func WelcomePlayer(text Choice, idPlayer int8) {
	// Welcome the player
	fmt.Println(WELCOME_PLAYER.Text)
	fmt.Printf(GIVE_PLAYER_ID.Text, idPlayer)
	// Ask for an opponent id
	reAskIdOpponent := true
	firstAsk := true
	for reAskIdOpponent {
		if firstAsk {
			idsOpponent = append(idsOpponent, askPlayer(ENTER_ID_OPPONENT.Text))
			firstAsk = false
		} else {
			reAskIdOpponent = askPlayer(ENTER_MORE_ID_OPPONENT.getTextWithChoices()) == COMMAND_YES
			if reAskIdOpponent {
				idsOpponent = append(idsOpponent, askPlayer(ENTER_ID_OPPONENT.Text))
			}
		}

	}
	fmt.Println(RULES)
}

func ActionMenu() {
	fmt.Println(PERSONNAL_ACTION_MENU.getTextWithChoices())
	fmt.Println(ACCESS_OPPONENT_ACTION_MENU.getTextWithChoices())
	playerChoice := askPlayer(ENTER_CHOICE.Text)
	switch playerChoice {
	case COMMAND_SEE_OWN_BOARD:
		showOwnBoard()
	case COMMAND_SEE_OWN_BOARD_STATE:
		showOwnBoardState()
	case COMMAND_ACTION_OPPONENT_BUTTON:
		OpponentActionMenu()
	default:
		fmt.Println(UNEXPECTED_ACTION.Text)
	}
	ActionMenu()
}

func OpponentActionMenu() {
	fmt.Println(COMBAT_MENU.Text)
	// idOpponent := askPlayer(WHICH_OPPONENT.getTextWithChoices())
	fmt.Println(OPPONENT_ACTION_MENU.getTextWithChoices())
	choicePlayer := askPlayer(ENTER_CHOICE.Text)
	switch choicePlayer {
	case COMMAND_ATTACK:
		// attack opponent
	case COMMAND_SEE_OPPONENT_BOARD:
		// see opponent board
	default:
		fmt.Println(UNEXPECTED_ACTION.Text)
	}
	ActionMenu()
}

// rajoute un joueur
func (choiceObj Choice) getTextWithChoices() string {
	choices := ""
	for _, choice := range choiceObj.Choices {
		choices += choice
	}
	return choiceObj.Text + choices
}

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
