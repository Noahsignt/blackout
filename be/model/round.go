package model

type Round struct {
	ID   string 			    `bson:"_id,omitempty" json:"id"`
	CurrRound int               `bson:"currRound" json:"currRound"`
	Bets      map[string]int    `bson:"bets" json:"bets"`                
	WonHands  map[string]int    `bson:"wonHands" json:"wonHands"`        
	Hands     map[string][]Card `bson:"hands" json:"hands"`     
	CurrHand  Hand              `bson:"currHand" json:"currHand"`
	Trump     int               `bson:"trump" json:"trump"`
}