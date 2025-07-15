package logic

import (
	"sync"
	"tiktok001/model"
)

type VideoServiceImpl struct {
	model.UserService
	model.CommentService
	model.LikeService
}

var (
	videoServiceImp  *VideoServiceImpl
	videoServiceOnce sync.Once
)
