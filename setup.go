package gongo

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

func (d *GongoClient) Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(d.URI))
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	d.Client = client
}

func (d *GongoClient) TestConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := d.Client.Ping(ctx, nil)
	if err != nil {
		fmt.Printf("%v\n\n", err)
		fmt.Println("Connection FAILED!")
	} else {
		fmt.Println("Connection SUCCESFULL!")
	}
}

func (d *GongoClient) GetAllDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() 
	dbs, err := d.Client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		fmt.Printf("Error findind all DB's")
	}
	fmt.Println(dbs)
}