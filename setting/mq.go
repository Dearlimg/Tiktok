package setting

import "tiktok001/setting/rabbitmq"

type mq struct {
}

func (m mq) Init() {
	rabbitmq.InitRabbitMQ()
	rabbitmq.InitLikeRabbitMQ()
	rabbitmq.InitFollowRabbitMQ()
	rabbitmq.InitCommentRabbitMQ()
}
