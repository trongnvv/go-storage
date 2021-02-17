package models

import "go.mongodb.org/mongo-driver/mongo"

var FileModel *mongo.Collection

type UserSchema struct {
	*BaseSchema `bson:",inline"`
	Path        string `bson:"path"`
	Name        string `bson:"name"`
	Type        string `bson:"type"`
}
