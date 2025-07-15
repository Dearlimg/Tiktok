package test

import (
	"fmt"
	"testing"
	"tiktok001/setting/rabbitmq"
)

func TestSimpleLikeMQPublish(t *testing.T) {
	rabbitmq.InitRabbitMQ()
	rabbitmq.InitLikeRabbitMQ()
	likeAddMQ := rabbitmq.SimpleLikeAddMQ
	userId := "2"
	videoId := "10"
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("%s-%s", userId, videoId)
		likeAddMQ.PublishSimple(msg)
	}
}

func TestSimpleCommentPublish(t *testing.T) {
	rabbitmq.InitRabbitMQ()
	rabbitmq.InitCommentRabbitMQ()
	commentDelMQ := rabbitmq.SimpleCommentDelMQ
	commentId := "26"
	for i := 0; i < 10; i++ {
		commentDelMQ.PublishSimple(commentId)
	}
}
