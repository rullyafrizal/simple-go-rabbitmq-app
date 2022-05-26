package configs

import (
	"fmt"
	"log"
	"simple-go-rabbitmq-app/utils"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitmqConfigContract interface {
	BuildConfigString() string
	ConnectToRabbitmq() *amqp091.Connection
}

type RabbitmqConfig struct {
	Host string
	Port string
	User string
	Pass string
}

func NewRabbitmqConfig() RabbitmqConfigContract {
	return &RabbitmqConfig{
		Host: utils.LookupEnv("RABBITMQ_HOST", "localhost"),
		Port: utils.LookupEnv("RABBITMQ_PORT", "5672"),
		User: utils.LookupEnv("RABBITMQ_USER", "guest"),
		Pass: utils.LookupEnv("RABBITMQ_PASS", "guest"),
	}
}

func (r *RabbitmqConfig) BuildConfigString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s/", r.User, r.Pass, r.Host, r.Port)
}

func (r *RabbitmqConfig) ConnectToRabbitmq() *amqp091.Connection {
	log.Println("Connecting to RabbitMQ...")
	conn, err := amqp091.Dial(r.BuildConfigString())
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ: ", err)
	}

	return conn
}
