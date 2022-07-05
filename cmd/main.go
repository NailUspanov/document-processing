package main

import (
	"github.com/sirupsen/logrus"
	"sstu/internal/handlers/rest"
	"sstu/internal/repos"
	mongo2 "sstu/internal/repos/mongo"
	"sstu/internal/services"
)

func main() {

	client, err := mongo2.ConnectDB()
	if err != nil {
		logrus.Fatal(err)
		return
	}
	collection := mongo2.GetCollection(client, "documents")
	repo := repos.NewRepository(collection)
	service := services.NewService(repo)
	handlers := rest.NewHandler(service)

	srv := new(Server)
	//TODO
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while runnning http server: %s", err.Error())
	}

}
