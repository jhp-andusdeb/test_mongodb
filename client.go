package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

func main() {
	Connect()
}

func Connect() {
	var err error

	var url string
	//host := "172.19.0.2:27017,172.19.0.3:27017,172.19.0.4:27017"
	//host := "172.19.0.3:27017"
	//dbOpt := "?replicaSet=replication"
	//url = fmt.Sprintf("mongodb://%s/%s", host, dbOpt)
	//MONGODB_URI="mongodb://localhost:27017,localhost:27018,localhost:27018/?replicaSet=rs1" TOPOLOGY=replica_set make

	//url = "mongodb://localhost:27020,localhost:27021,localhost:27022/admin?replSet=replication"
	url = "mongodb://localhost:27020,localhost:27021,localhost:27022/admin?replSet=replication"

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		fmt.Println("failed", url)
	}

	err = client.Connect(context.Background())
	if err != nil {
		fmt.Println("connect fail", err)
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		fmt.Println("ping failed!", err)
		os.Exit(1)
	} else {
		fmt.Println("Successfully Connected to MongoDB!")
	}

	//test := client.Database("anduschain").Collection("test")

	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("successfully disconnected")
}

func Start() error {


	//logger.Debug("Start fairnode mongo database", "chainID", chainID.String(), "url", url)

	return nil
}

func Stop() {

}
