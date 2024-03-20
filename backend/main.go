package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/K1ose/bdsms/backend/controllers"
	"github.com/K1ose/bdsms/backend/repositories"
	"github.com/K1ose/bdsms/backend/services"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigName("config") // 配置文件名（无扩展名）
	viper.SetConfigType("yaml")   // 配置文件类型，例如 YAML
	viper.AddConfigPath(".")      // 配置文件路径

	// 尝试读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}

func main() {
	// 数据库连接字符串
	initConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("database.host"),
		viper.GetString("database.user"),
		viper.GetString("database.dbname"),
		viper.GetString("database.password"),
		viper.GetString("database.sslmode"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 初始化仓库
	userRepo := repositories.NewPostgresUserRepository(db)

	// 初始化服务
	userService := services.NewUserService(userRepo)

	// 初始化控制器
	userController := controllers.NewUserController(userService)

	// 初始化 Gin 路由
	router := gin.Default()

	// 使用 SetUpRoutes 函数设置路由
	SetUpRoutes(router, userController)

	// 启动 Gin 服务器
	router.Run(":8080")
}
