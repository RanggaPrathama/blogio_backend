package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tags struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	NAME_TAGS string             `json:"name_tags,omitempty" bson:"name_tags,omitempty" validate:"required"`
}
