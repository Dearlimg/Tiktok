package setting

type group struct {
	DB database
	MQ mq
}

var Group = new(group)

func Init() {
	Group.DB.Init()
	Group.MQ.Init()
}
