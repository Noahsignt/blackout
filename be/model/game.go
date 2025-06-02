package model

type Game struct {
    ID   string `bson:"_id,omitempty" json:"id"`
    NumRounds int `bson:"numRounds" json:"numRounds"`
    Round Round `bson:"round" json:"round"`
    Players []Player `bson:"players" json:"players"`
}