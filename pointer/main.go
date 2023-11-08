package main

import "fmt"

type Player struct {
	playerHealth int
}

// IF THE player IS NOT a pointer we are ADJUSTing just a copy of the player in the function
// Not actual player......
//
//	func takeDamage(player *Player, amount int) {
//		player.playerHealth -= amount
//		fmt.Println("The player is taking damage..New health is ->", player.playerHealth)
//	}
func (player *Player) takeDamage(amount int) {
	player.playerHealth -= amount
	fmt.Println("The player is taking damage..New health is ->", player.playerHealth)
}

func main() {
	player := &Player{
		playerHealth: 100,
	}

	player.takeDamage(10)

	//we see health is 90 but after the function playerHealth is 100
	fmt.Println(player)
}
