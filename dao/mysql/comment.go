package mysql

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm" // 需导入 gorm
)

type Comment struct {
	Id         int64     // 评论id
	UserId     int64     // 评论用户id
	VideoId    int64     // 视频id
	Content    string    // 评论内容
	ActionType int64     // 1:发布，2:删除
	CreatedAt  time.Time // 发布时间
	UpdatedAt  time.Time
}

func (Comment) TableName() string {
	return "comment"
}

// 用指针接收器，确保外部实例同步更新
func (comment *Comment) InsertComment() (*Comment, error) {
	if err := Db.Model(Comment{}).Create(comment).Error; err != nil {
		log.Printf("插入评论失败: %v", err)
		return nil, err
	}
	return comment, nil
}

// 指针接收器，优化错误处理和局部变量
func (comment *Comment) DeleteComment(commentId int64) error {
	var temp Comment
	result := Db.Where("id = ?", commentId).First(&temp)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("评论不存在")
		}
		log.Printf("查询评论失败（id=%d）: %v", commentId, result.Error)
		return errors.New("查询评论失败")
	}
	// 更新action_type为2（逻辑删除）
	result = Db.Model(Comment{}).Where("id=?", commentId).Update("action_type", 2)
	if result.Error != nil {
		log.Printf("删除评论失败（id=%d）: %v", commentId, result.Error)
		return errors.New("删除评论失败")
	}
	return nil
}

// 改为包级函数，无需依赖实例
func GetCommentList(videoId int64) ([]Comment, error) {
	var commentList []Comment
	result := Db.Model(Comment{}).Where(map[string]interface{}{"video_id": videoId, "action_type": 1}).
		Order("created_at desc").Find(&commentList)
	if result.Error != nil {
		log.Printf("查询视频评论列表失败（videoId=%d）: %v", videoId, result.Error)
		return commentList, errors.New("get comment list failed")
	}
	return commentList, nil
}

// 改为包级函数
func GetCommentCnt(videoId int64) (int64, error) {
	var count int64
	result := Db.Model(Comment{}).Where(map[string]interface{}{"video_id": videoId, "action_type": 1}).Count(&count)
	if result.Error != nil {
		log.Printf("查询视频评论数量失败（videoId=%d）: %v", videoId, result.Error)
		return 0, errors.New("find comments count failed")
	}
	return count, nil
}
