package main

import (
	"fmt"
	"log"
	"net/http"

	"card-system/internal/config"
	"card-system/internal/controller"
	"card-system/internal/middleware"
	"card-system/internal/model"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("../config")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	err = model.ConnectDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// 自动迁移模型
	model.DB.AutoMigrate(
		&model.User{},
		&model.Merchant{},
		&model.Card{},
		&model.Order{},
	)

	// 设置路由
	r := gin.Default()

	// 注册控制器
	userCtrl := controller.NewUserController(cfg)
	merchantCtrl := controller.NewMerchantController(cfg)

	// 公共路由
	public := r.Group("/api")
	{
		public.POST("/register", userCtrl.Register)
		public.POST("/login", userCtrl.Login)
	}

	// 认证路由
	authGroup := r.Group("/api")
	authGroup.Use(middleware.AuthMiddleware(cfg))
	{
		authGroup.GET("/user/me", userCtrl.GetCurrentUser)
		authGroup.POST("/merchants", merchantCtrl.CreateMerchant)
	}

	// 管理员路由
	adminGroup := r.Group("/api/admin")
	adminGroup.Use(middleware.AuthMiddleware(cfg), middleware.AdminMiddleware())
	{
		adminGroup.GET("/merchants", merchantCtrl.GetAllMerchants)
		adminGroup.PUT("/merchants/:id/status", merchantCtrl.UpdateMerchantStatus)
	}

	// 启动服务器
	log.Printf("Server started on port %s", cfg.App.Port)
	log.Fatal(http.ListenAndServe(cfg.App.Port, r))
}
