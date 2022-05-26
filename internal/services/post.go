package services

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"simple-go-rabbitmq-app/internal/models"
	"strings"
	"sync"

	"github.com/rabbitmq/amqp091-go"
)

type PostServiceContract interface {
	PublishPost(post models.Post) error
	ConsumePost() error
}

type PostService struct {
	channel *amqp091.Channel
}

func NewPostService(ch *amqp091.Channel) PostServiceContract {
	return &PostService{
		channel: ch,
	}
}

func (p *PostService) PublishPost(post models.Post) error {
	err := p.channel.ExchangeDeclare(
		"post-exchange-direct", // name
		"direct",
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	body, err := json.Marshal(post)
	if err != nil {
		return err
	}

	log.Println("Publishing post to RabbitMQ...")

	err = p.channel.Publish(
		"post-exchange-direct",              // exchange
		determineRoutingKey(int64(post.ID)), // routing key
		false,                               // mandatory
		false,                               // immediate
		amqp091.Publishing{
			DeliveryMode: amqp091.Persistent, // the message will survive server restarts/down
			ContentType:  "application/json",
			Body:         body,
		},
	)
	if err != nil {
		return err
	}

	log.Println("Post successfully published to RabbitMQ")

	return nil
}

func (p *PostService) ConsumePost() error {
	err := p.channel.ExchangeDeclare(
		"post-exchange-direct", // name
		"direct",
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	queue, err := p.channel.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = p.channel.QueueBind(
		queue.Name,                        // queue name
		strings.Split(os.Args[1], "=")[1], // routing key
		"post-exchange-direct",            // exchange
		false,                             // no-wait
		nil,                               // arguments
	)
	if err != nil {
		return err
	}

	post, err := p.channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return err
	}

	log.Println("Consuming posts from RabbitMQ...")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		// this is where we iterate over the channel
		// for range will wait and stand by to display post until the channel is closed
		for p := range post {
			displayPosts(p.Body)
		}
	}()

	wg.Wait()

	return err
}

func displayPosts(body []byte) {
	var post models.Post

	err := json.Unmarshal(body, &post)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("")
	log.Println("ID:", post.ID)
	log.Println("Title: ", post.Title)
	log.Println("Content: ", post.Content)
	log.Println("Image: ", post.Image)
	log.Println("CreatedAt: ", post.CreatedAt)
	log.Println("UpdatedAt: ", post.UpdatedAt)
	fmt.Println("")
}

func determineRoutingKey(id int64) string {
	if id%3 == 0 && id%5 == 0 {
		return "fizzbuzz"
	} else if id%3 == 0 {
		return "fizz"
	}

	return "buzz"
}
