package logic

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"tiktok001/config"
	"tiktok001/dao/mysql"
	"tiktok001/model"
	"time"
)

type MessageServiceImpl struct {
}

var (
	messageServiceImpl *MessageServiceImpl
	messageServiceOnce sync.Once
)

// GetMessageServiceInstance Go 单例模式
func GetMessageServiceInstance() *MessageServiceImpl {
	messageServiceOnce.Do(func() {
		messageServiceImpl = &MessageServiceImpl{}
	})
	return messageServiceImpl
}

func (messageService *MessageServiceImpl) SendMessage(fromUserId int64, toUserId int64, content string, actionType int64) error {
	var err error
	switch actionType {
	// actionType = 1 发送消息
	case 1:
		err = mysql.SendMessage(fromUserId, toUserId, content, actionType)
	default:
		log.Println(fmt.Sprintf("未定义 actionType=%d", actionType))
		return errors.New(fmt.Sprintf("未定义 actionType=%d", actionType))
	}
	return err
}

func (messageService *MessageServiceImpl) MessageChat(loginUserId int64, targetUserId int64, latestTime time.Time) ([]model.Message, error) {
	messages := make([]model.Message, 0, config.VIDEO_INIT_NUM_PER_AUTHOR)
	plainMessages, err := mysql.MessageChat(loginUserId, targetUserId, latestTime)
	if err != nil {
		log.Println("MessageChat Service:", err)
		return nil, err
	}
	err = messageService.getRespMessage(&messages, &plainMessages)
	if err != nil {
		log.Println("getRespMessage:", err)
		return nil, err
	}
	return messages, nil
}

func (messageService *MessageServiceImpl) LatestMessage(loginUserId int64, targetUserId int64) (model.LatestMessage, error) {
	plainMessage, err := mysql.LatestMessage(loginUserId, targetUserId)
	if err != nil {
		log.Println("LatestMessage Service:", err)
		return model.LatestMessage{}, err
	}
	var latestMessage model.LatestMessage
	latestMessage.Message = plainMessage.MsgContent
	if plainMessage.UserId == loginUserId {
		// 最新一条消息是当前登录用户发送的
		latestMessage.MsgType = 1
	} else {
		// 最新一条消息是当前好友发送的
		latestMessage.MsgType = 0
	}
	return latestMessage, nil
}

// 返回 message list 接口所需的 Message 结构体
func (messageService *MessageServiceImpl) getRespMessage(messages *[]model.Message, plainMessages *[]mysql.Message) error {
	for _, tmpMessage := range *plainMessages {
		var message model.Message
		messageService.combineMessage(&message, &tmpMessage)
		*messages = append(*messages, message)
	}
	return nil
}

func (messageService *MessageServiceImpl) combineMessage(message *model.Message, plainMessage *mysql.Message) error {
	message.Id = plainMessage.Id
	message.UserId = plainMessage.UserId
	message.ReceiverId = plainMessage.ReceiverId
	message.MsgContent = plainMessage.MsgContent
	message.CreatedAt = plainMessage.CreatedAt.Unix()
	return nil
}
