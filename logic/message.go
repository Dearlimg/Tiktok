package logic

import "sync"

type MessageServiceImpl struct {
}

var (
	messageServiceImpl *MessageServiceImpl
	messageServiceOnce sync.Once
)
