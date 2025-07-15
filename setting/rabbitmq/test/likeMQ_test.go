package test

import (
	"testing"
	"tiktok001/setting/rabbitmq"
)

func TestSimpleLikeMQConsumer(t *testing.T) {
	rabbitmq.InitRabbitMQ()
	rabbitmq.InitLikeRabbitMQ()
	likeAddMQ := rabbitmq.SimpleLikeAddMQ
	likeAddMQ.ConsumeSimple()
}

func TestSimpleCommentConsumer(t *testing.T) {
	rabbitmq.InitRabbitMQ()
	rabbitmq.InitCommentRabbitMQ()
	commentDelMQ := rabbitmq.SimpleCommentDelMQ
	commentDelMQ.ConsumeSimple()
}
