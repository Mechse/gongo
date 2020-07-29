package gongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type GongoClient struct {
	URI    string
	Client *mongo.Client
}

