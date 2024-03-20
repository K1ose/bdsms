package controllers

import (
	"net/http"

	"github.com/K1ose/bdsms/backend/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

// 假设 YourUserServiceType 是定义在 services 包中的类型，且它有一个 RegisterUser 方法
// 请根据实际情况调整 YourUserServiceType 和 RegisterUser 方法的具体实现

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (c *UserController) Register(ctx *gin.Context) {
	var input struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"` // 注意：实际应用中应先进行加密处理
		CreatedAt string `json:"created_at"`
	}

	// 使用 ctx.ShouldBindJSON 来解析请求体中的 JSON 数据到 input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用 userService 的 RegisterUser 方法来注册用户
	// 密码加密操作应该在 userService.RegisterUser 方法内部完成
	if err := c.userService.RegisterUser(input.Username, input.Email, input.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	// 发送成功响应
	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
