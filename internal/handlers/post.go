package handlers

import (
	"simple-go-rabbitmq-app/internal/models"
	"simple-go-rabbitmq-app/internal/services"
	"strconv"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type PostHandlerContract interface {
	CreatePost() error
	GetPosts() error
}

type PostHandler struct {
	Service services.PostServiceContract
}

func NewPosthandler(ch *amqp091.Channel) PostHandlerContract {
	return &PostHandler{
		Service: services.NewPostService(ch),
	}
}

func (h *PostHandler) CreatePost() error {
	for i := 0; i < 10; i++ {
		post := models.Post{
			ID:        i,
			Title:     "This is post " + strconv.Itoa(i),
			Content:   "This is content " + strconv.Itoa(i),
			Image:     "This is image " + strconv.Itoa(i),
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		err := h.Service.PublishPost(post)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *PostHandler) GetPosts() error {
	err := h.Service.ConsumePost()
	if err != nil {
		return err
	}

	return nil
}
