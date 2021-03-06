package services

import (
	"encoding/json"
	"log"
	"simple-go-rabbitmq-app/internal/models"

	"github.com/rabbitmq/amqp091-go"
)

type PostServiceContract interface {
	PublishPost(post models.Post) error
	ConsumePost() ([][]byte, error)
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
	queue, err := p.channel.QueueDeclare(
		"post-queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
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
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,		// immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	log.Println("Post successfully published to RabbitMQ")

	return nil
}

func (p *PostService) ConsumePost() ([][]byte, error) {
	var posts [][]byte

	queue, err := p.channel.QueueDeclare(
		"post-queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return posts, err
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
		return posts, err
	}

	log.Println("Consuming posts from RabbitMQ...")

	// this is where we iterate over the channel
	// for range will wait and stand by to display post until the channel is closed
	for p := range post {
		displayPosts(p.Body)
	}

	defer log.Println("Posts successfully consumed from RabbitMQ")

	return posts, nil
}

func displayPosts(body []byte) {
	var post models.Post

	err := json.Unmarshal(body, &post)
	if err != nil {
		log.Println(err)
	}

	log.Println("===========================")
	log.Println("Title: ", post.Title)
	log.Println("Content: ", post.Content)
	log.Println("Image: ", post.Image)
	log.Println("CreatedAt: ", post.CreatedAt)
	log.Println("UpdatedAt: ", post.UpdatedAt)
	log.Println("===========================")
}

func printPost(ch <-chan amqp091.Delivery) {
	for d := range ch {
		var post models.Post

		err := json.Unmarshal(d.Body, &post)
		if err != nil {
			log.Println(err)
		}

		log.Println("===========================")
		log.Println("Title: ", post.Title)
		log.Println("Content: ", post.Content)
		log.Println("Image: ", post.Image)
		log.Println("CreatedAt: ", post.CreatedAt)
		log.Println("UpdatedAt: ", post.UpdatedAt)
		log.Println("===========================")
	}
}
