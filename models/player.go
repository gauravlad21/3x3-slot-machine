package models

type Player struct {
	ID      string `json:"id" bson:"_id"`
	Name    string `json:"name" bson:"name"`
	Credits int    `json:"credits" bson:"credits"`
	Status  string `json:"status" bson:"status"`
}
