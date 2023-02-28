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
	strId,_ := strconv.Atoi(id)

	response, err := http.Get(baseUrl+"users/get?id="+strId)
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