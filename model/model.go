package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID      primitive.ObjectID `json:"_id"`
	NAME    string             `json:"name"`
	EMAIL   string             `json:"email"`
	PHONENO string             `json:"phoneno"`
}
