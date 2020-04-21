package controller

import (
	"gim/internal/logic/model"
	"gim/internal/logic/service"
)

func init() {
	g := Engine.Group("/user")
	g.POST("", handler(UserController{}.Add))
	g.GET("/info", handler(UserController{}.UserInfo))
}

type UserController struct{}

// Add 添加用户
func (UserController) Add(c *context) {
	var user model.User
	if c.bindJson(&user) != nil {
		return
	}
	user.AppId = c.appId
	c.response(nil, service.UserService.Add(c, user))
}

// UserInfo 获取用户信息
func (UserController) UserInfo(c *context) {
	c.response(service.UserService.Get(c, c.appId, c.userId))
}
