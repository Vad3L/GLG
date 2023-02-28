package main

type Planets struct {
	OwnerId	string
	HqLevel	int
}

type Player struct {
	Id	string
	Name	string
	Avatar	string
	Level	int
	Experience	int
	AlianceId	string
	Planets	[]*Planets
}

type Server struct {
	Name	string
	IsOnline	bool
	Ping	int
}