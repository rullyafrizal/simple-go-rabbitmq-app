package test

import (
	"log"
	"simple-go-rabbitmq-app/configs"
	"simple-go-rabbitmq-app/internal/handlers"
	"testing"

	"github.com/joho/godotenv"
)

func TestPublish(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	rbmqConf := configs.NewRabbitmqConfig()
	conn := rbmqConf.ConnectToRabbitmq()
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel: ", err)
	}
	defer ch.Close()

	postHandler := handlers.NewPosthandler(ch)

	err = postHandler.CreatePost()
	if err != nil {
		log.Fatal("Failed to create post: ", err)
	}
}
