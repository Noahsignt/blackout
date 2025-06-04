package model

type Hand struct {
    WinningCard   CardWithPlayer `bson:"winningCard" json:"winningCard"`
    StartingIdx   int            `bson:"startingIdx" json:"startingIdx"`
    CurrPlayer    int            `bson:"currPlayer" json:"currPlayer"`
}

type CardWithPlayer struct {
    Card   Card   `bson:"card" json:"card"`
    Player Player `bson:"player" json:"player"`
}

func NewHand(startingIdx int) Hand{
	return Hand{StartingIdx: startingIdx, CurrPlayer: 0}
}