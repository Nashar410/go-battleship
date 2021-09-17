package main

var idPlayer = RandomString(4)
var idsOpponent []string
var caseTouched []ShipPosition
var ships []Ship

func main() {
	StartServer()
	ships := generateShips()
	GenerateAndShowABoard(ships)
	WelcomePlayer(idPlayer)
	ActionMenu()

}

func generateShips() []Ship {
	// Add Two boats
	ships = append(ships, generateShip(ships, 2))
	ships = append(ships, generateShip(ships, 3))
	ships = append(ships, generateShip(ships, 4))
	return ships
}
