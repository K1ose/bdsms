// router.go
package main

import (
	"github.com/K1ose/bdsms/backend/controllers"
	"github.com/gin-gonic/gin"
)

// SetUpRoutes 设置应用的所有路由
func SetUpRoutes(router *gin.Engine, userController *controllers.UserController) {
	// 这里注册所有的路由和对应的控制器方法
	router.POST("/register", userController.Register)
	// 可以继续添加其他路由
}
