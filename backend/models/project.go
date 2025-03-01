package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Project struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Title       string             `bson:"title"`
    Description string             `bson:"description"`
    Link        string             `bson:"link"`
}
