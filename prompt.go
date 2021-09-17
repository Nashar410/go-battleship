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
	[]string{
		COMMAND_YES + " pour oui\n",
		COMMAND_NO + " pour non\n"}}

var RULES = Choice{
	"",
	[]string{}}

var PERSONNAL_ACTION_MENU = Choice{
	"Voici les actions possibles: ",
	[]string{
		COMMAND_SEE_OWN_BOARD + " pour voir la position de vos bateaux\n",
		COMMAND_SEE_OWN_BOARD_STATE + " pour voir l'état de votre plateau\n"}}

var ACCESS_OPPONENT_ACTION_MENU = Choice{
	"Accéder au menu de combat\n",
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

var ATTACKED_SUCCESSED= Choice{
	"Un de vos navires a été touché par un tir de %s\n",
	[]string{}}

var ATTACKED_FAILED= Choice{
	"Vous avez essuyé un tir de %s sans dégats %s\n",
	[]string{}}

var ATTACKED_SINKED= Choice{
	"Vous n'avez touché aucun des navires de %s, un de vos navires a coulé\n",
	[]string{}}

var YOU_LOST = Choice{
	"Tout vos navires ont été coulés, vous avez perdu.\n",
	[]string{}}


var OPPONENT_LOST = Choice{
	"%s n'a plus aucun navire, il a perdu.\n",
	[]string{}}

var YOU_WIN  = Choice{
	"Vous êtes le dernier en lice ! Vous avez gagné !\n",
	[]string{}}

type Choice struct {
	Text string
	Choices []string
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