package example

import (
	"fmt"
	."github.com/vad3l/galaxylifegoapi"
)



func main() {
	player := GetUserByName("vadeledav")
	playerList := SearchUserByName("vad")
	serversList := GetServerStatus()
	playerId := GetUserById(268172)
	userStats := GetUserStats(268172)
	alliance := GetAlliance("France Leader")

	fmt.Println("GetUserByName : \n",player)
	fmt.Println(player.Planets[1],"\n")

	fmt.Println("SearchUserByName : \n",playerList[1],"\n")
	fmt.Println("GetServersStatus : \n",serversList,"\n")
	fmt.Println("GetUserById : \n",playerId,"\n")
	fmt.Println("GetUserStats : \n",userStats,"\n")
	fmt.Println("GetAlliance : \n",alliance)
	fmt.Println(alliance.Members[0],"\n")
}