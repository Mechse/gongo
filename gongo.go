package gongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GongoClient struct {
	URI    string
	Client *mongo.Client
}

func (d *GongoClient) Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(d.URI))
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	d.Client = client
}
