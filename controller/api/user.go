package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok001/dao/mysql"
	"tiktok001/logic"
	"tiktok001/model"
	"tiktok001/utils"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User model.User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	usi := logic.GetUserServiceInstance()
	user := usi.GetUserBasicInfoByName(username)
	if username == user.Name {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		newUser := mysql.UserBasicInfo{
			Name:     username,
			Password: logic.EnCoder(password),
		}
		if usi.InsertUser(&newUser) != true {
			fmt.Println("Insert Fail")
		}
		// 得到用户id
		user := usi.GetUserBasicInfoByName(username)
		userId := user.Id
		token := utils.GenerateToken(userId, username)
		//log.Println("注册时返回的token", token)
		//log.Println("注册返回的id: ", user.Id)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Register Success"},
			UserId:   user.Id,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	encoderPassword := logic.EnCoder(password)
	//log.Println("encoderPassword:", encoderPassword)
	// 登录逻辑：使用jwt，根据用户信息生成token
	usi := logic.GetUserServiceInstance()

	user := usi.GetUserBasicInfoByName(username)
	userId := user.Id
	if encoderPassword == user.Password {
		token := utils.GenerateToken(userId, username)
		//log.Println("generate token:", token)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Login Success"},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User or Password Error"},
		})
	}
}

func UserInfo(c *gin.Context) {
	userId := c.Query("user_id")
	// 使用中间件，做token权限校验
	usi := logic.GetUserServiceInstance()
	id, _ := strconv.ParseInt(userId, 10, 64)
	if user, err := usi.GetUserLoginInfoById(id); err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Query Success"},
			User:     user,
		})
	}
}

func Test(c *gin.Context) {
	// 通过c.Get()获取userId
	userId, _ := c.Get("userId")
	//fmt.Println("userId", userId)
	c.JSON(http.StatusOK, userId)
}
