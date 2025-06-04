package model

import (
	"math/rand"
)

type Round struct {
	RoundNum int               	`bson:"roundNum" json:"roundNum"`
	Bets      map[string]int    `bson:"bets" json:"bets"`                
	WonHands  map[string]int    `bson:"wonHands" json:"wonHands"`        
	Hands     map[string][]Card `bson:"hands" json:"hands"`     
	CurrHand  Hand              `bson:"currHand" json:"currHand"`
	Trump     int               `bson:"trump" json:"trump"`
}

func NewRound(roundNum int, currHand Hand) *Round {
	trump := rand.Intn(4) + 1

    return &Round{
        RoundNum: roundNum,
        Bets:     make(map[string]int),
        WonHands: make(map[string]int),
        Hands:    make(map[string][]Card),
        CurrHand: currHand,
        Trump:    trump,
    }
}
