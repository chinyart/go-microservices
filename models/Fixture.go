package models

type Fixture struct {
	Team1 string `json: "team1" bson: "team"`
	Team2 string
	score int
	date  string
	vanue string
}
