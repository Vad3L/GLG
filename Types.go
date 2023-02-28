package galaxyLifeAPI

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

type UserStats struct {
	TotalPlayTimeInMs	int64
	NpcsAttacked	int
	PlayersAttacked	int
	TimesAttacked	int
	StarbasesDestroyed	int
	BuildingsDestroyed	int
	DamageDoneInAttacks	int
	ObstaclesRecycled	int
	CoinsSpent	int64
	MineralsSpent	int64
	CoinsFromAttacks	int64
	MineralsFromAttacks	int64
	ScoreFromAttacks	int64
	NukesUsed	int
	TroopsTrained	int
	TroopSizesDonated	int
	FriendsHelped	int
	GiftsReceived	int
	GiftsSent	int
	ColoniesMoved	int
	RivalsWon	int
	ChipsWonFromRivals	int
}

type Alliance struct {
	Id string
	Name	string
	Description	string
	Emblem	EmblemStruct
	AllianceLevel	int
	WarPoints	int64
	WarsWon	int
	WarsLost	int
	Members	[]*AllianceMember
}

type EmblemStruct struct {
	Shape	int
	Pattern	int
	Icon	int
}

type AllianceMember struct {
	Id	int
	Name	string
	Avatar	string
	AllianceRole	int
	TotalWarPoints	int
}