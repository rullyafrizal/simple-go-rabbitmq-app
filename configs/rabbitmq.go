package configs

import (
	"fmt"
	"simple-go-rabbitmq-app/utils"
)

type RabbitmqConfigContract interface {
	BuildConfigString() string
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