package config

import "time"

// 存放相关配置

const GO_STARTER_TIME = "2006-01-02 15:04:05"

// 视频模块相关配置
const (
	VIDEO_NUM_PER_REFRESH     = 6
	VIDEO_INIT_NUM_PER_AUTHOR = 10
	// 阿里 OSS 相关配置
	OSS_ACCESS_KEY_ID     = "LTAI5t7Yy1jQVLAvVAHF2KqZ"
	OSS_ACCESS_KEY_SECRET = "6MjwdHPszGfOnBJBJklXzjqTcEQRlO"
	OSS_BUCKET_NAME       = "simple-tiktok"
	OSS_ENDPOINT          = "https://oss-cn-beijing.aliyuncs.com"
	CUSTOM_DOMAIN         = "123.249.42.125/"
	OSS_VIDEO_DIR         = "tiktok/"
	PLAY_URL_PREFIX       = CUSTOM_DOMAIN + OSS_VIDEO_DIR
	COVER_URL_SUFFIX      = "?x-oss-process=video/snapshot,t_2000,m_fast"
	OSS_USER_AVATAR_DIR   = ""
)

const LIKE = 1

const BG_IMAGE = "nil"

const SIGNATURE = "nil"

const SECRETE = "joker"

const (
	ExpireTime = 100000
)

var LatestRequestTime = make(map[string]time.Time, 100)
