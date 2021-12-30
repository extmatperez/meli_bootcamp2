package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MongoDB *mongo.Client
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	MongoDB, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Si lo uso, desconecta la base de datos, sino funca bien pero debe quedar abierta la conexion
	// defer func() {
	// 	if err = MongoDB.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := MongoDB.Ping(ctx, readpref.Primary()); err != nil {
		// Can't connect to Mongo server
		log.Fatal(err)
	}
}
