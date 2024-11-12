package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	USERNAME string `json:"username,omitempty" bson:"username,omitempty"`
	EMAIL string `json:"email,omitempty" bson:"email,omitempty"`
	PASSWORD string `json:"password,omitempty" bson:"password,omitempty"`
	JENIS_KELAMIN bool `json:"jenis_kelamin,omitempty" bson:"jenis_kelamin,omitempty"`
	PHOTO string `json:"photo,omitempty" bson:"photo,omitempty"`
	PHONE string `json:"phone,omitempty" bson:"phone,omitempty"`
}