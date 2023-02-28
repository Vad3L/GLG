package main

import (
	"fmt"
)



func main() {
	player := GetUserByName("vadeledav")
	playerList := SearchUserByName("vad")
	serversList := GetServerStatus()
	playerId := GetUserById(268172)

	fmt.Println(player)
	fmt.Println(player.Planets[1])

	fmt.Println(playerList[1])

	fmt.Println(serversList)
	fmt.Println(playerId)

}