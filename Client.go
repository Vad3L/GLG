package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"strconv"
)

var baseUrl string = "https://api.galaxylifegame.net/"

func GetServerStatus() ([]Server) {
	response, err := http.Get(baseUrl+"status")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	var servers []Server 
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {log.Fatal(err)}

	err = json.Unmarshal(responseData, &servers)
	if err != nil {fmt.Println(err)}

	return servers
}

func GetUserByName(name string) (Player) {
	
	response, err := http.Get(baseUrl+"users/name?name="+name)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	var player Player 
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {log.Fatal(err)}
	err = json.Unmarshal(responseData, &player)
	if err != nil {fmt.Println(err)}

	return player
}


func GetUserById(id int) (Player) {
	response, err := http.Get(baseUrl+"users/get?id="+strconv.Itoa(id))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	var player Player 
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {log.Fatal(err)}

	err = json.Unmarshal(responseData, &player)
	if err != nil {fmt.Println(err)}

	return player
}

func SearchUserByName(name string) ([]Player) {
	response, err := http.Get(baseUrl+"users/search?name="+name)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	var player []Player 
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {log.Fatal(err)}

	err = json.Unmarshal(responseData, &player)
	if err != nil {fmt.Println(err)}

	return player
}

func GetUserBySteamId(id int) (Player) {
	response, err := http.Get(baseUrl+"users/steam?steamId="+strconv.Itoa(id))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	var player Player 
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {log.Fatal(err)}

	err = json.Unmarshal(responseData, &player)
	if err != nil {fmt.Println(err)}

	return player
}

func GetUserStats(id int) (UserStats) {
	response, err := http.Get(baseUrl+"users/stats?id="+strconv.Itoa(id))
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	var stats UserStats 
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {log.Fatal(err)}

	err = json.Unmarshal(responseData, &stats)
	if err != nil {fmt.Println(err)}

	return stats
}

func GetAlliance(id string) (Alliance) {
	response, err := http.Get(baseUrl+"alliances/get?name="+id)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	var alliance Alliance 
	
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {log.Fatal(err)}

	err = json.Unmarshal(responseData, &alliance)
	if err != nil {fmt.Println(err)}

	return alliance
}