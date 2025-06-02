package model

type Player struct {
	ID string			    `bson:"_id,omitempty" json:"id"`
	Score int               `bson:"score" json:"score"`
	Name string    			`bson:"name" json:"name"`
}