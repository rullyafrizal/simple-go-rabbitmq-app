package main

import (
	"log"
	"simple-go-rabbitmq-app/configs"
	"simple-go-rabbitmq-app/internal/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
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

	err = postHandler.GetPosts()
	if err != nil {
		log.Fatal("Failed to get posts: ", err)
	}
}
