package mysql

import (
	"errors"
	"log"
	"time"
)

type Like struct {
	Id        int64
	UserId    int64
	VideoId   int64
	Liked     int8
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (l Like) TableName() string {
	return "like"
}

// GetLikeListByUserId 获取当前用户点赞视频id列表
func GetLikeListByUserId(userId int64) ([]int64, int64, error) {
	var LikedList []int64
	result := Db.Model(&Like{}).Where("user_id=? and liked=?", userId, 1).Order("created_at desc").Pluck("video_id", &LikedList)
	likeCnt := result.RowsAffected
	if result.Error != nil {
		log.Println("LikedVideoIdList:", result.Error.Error())
		return nil, -1, result.Error
	}
	return LikedList, likeCnt, nil
}

// VideoLikedCount 统计视频点赞数量
func VideoLikedCount(videoId int64) (int64, error) {
	var count int64
	//数据库中查询点赞数量
	err := Db.Model(Like{}).Where(map[string]interface{}{"video_id": videoId, "liked": 1}).Count(&count).Error
	if err != nil {
		log.Println("LikeDao-Count: return count failed") //函数返回提示错误信息
		return -1, errors.New("find likes count failed")
	}
	log.Println("LikeDao-Count: return count success") //函数执行成功，返回正确信息
	return count, nil
}
