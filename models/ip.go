package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// IP struct
type IP struct {
	ID         bson.ObjectId `bson:"_id" json:"-"`
	Data       string        `bson:"data" json:"ip"`
	Type       string        `bson:"type" json:"type"`
	CreateTime time.Time     `bson:"create_time" json:"create_time"`
	UpdateTime time.Time     `bson:"update_time" json:"update_time"`
}

// NewIP .
func NewIP() *IP {
	return &IP{
		ID: bson.NewObjectId(),
	}
}
