package models

import "time"

type Customer struct {
	Id        uint32    `bson:"_id" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}
