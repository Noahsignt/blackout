package model

type Hand struct {
    WinningCard   CardWithPlayer `bson:"winningCard" json:"winningCard"`
    SortedPlayers []Player       `bson:"sortedPlayers" json:"sortedPlayers"`
    CurrPlayer    int         `bson:"currPlayer" json:"currPlayer"`
}

type CardWithPlayer struct {
    Card   Card   `bson:"card" json:"card"`
    Player Player `bson:"player" json:"player"`
}
