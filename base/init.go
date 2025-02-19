package base

import (
	"backend/constants"
	"backend/env"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBInstance *mongo.Client

func ConnectDB() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer ctxCancel()

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	UriEnv := env.Get(constants.MONGODBURI)
	option := options.Client().ApplyURI(UriEnv).SetServerAPIOptions(serverAPI)

	clients, err := mongo.Connect(ctx, option)
	if err != nil {
		fmt.Println(err)
	}

	DBInstance = clients
}
