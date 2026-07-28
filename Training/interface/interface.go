package main

import "fmt"

type runner interface {
	Run()
}

type shooter interface {
	shoot()
}

type walker interface {
	Walk()
}

type player struct {
	name       string
	age        int
	position   string
	strongFoot string
}

func main() {

	player1 := &player{
		name:       "amirreza",
		age:        22,
		position:   "benchWarmer",
		strongFoot: "noFoot",
	}

	var Runner runner = player1
	var Walker walker = player1

	Runner.Run()
	Walker.Walk()

}

func (player *player) Run() {
	fmt.Printf("name: %s, position: %s, player is Running\n", player.name, player.position)
}
func (player *player) Walk() {
	fmt.Printf("name: %s, position: %s, player is walking\n", player.name, player.position)
}
