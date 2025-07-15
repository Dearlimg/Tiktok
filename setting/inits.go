package setting

import "tiktok001/setting/rabbitmq"

type group struct {
	DB database
}

var Group = new(group)

func Init() {
	Group.DB.Init()
	rabbitmq.InitRabbitMQ()
	rabbitmq.InitLikeRabbitMQ()
	rabbitmq.InitFollowRabbitMQ()
	rabbitmq.InitCommentRabbitMQ()
}
