// cmd/main.go
package main

import (
	"card-system/internal/config"
	"card-system/internal/controller"
	"card-system/internal/middleware"
	"card-system/internal/model"
	"card-system/internal/repository"
	"card-system/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("../config")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 连接数据库
	dsn := cfg.DB.DSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// 自动迁移模型
	db.AutoMigrate(
		&model.User{},
		&model.Merchant{},
		&model.Page{},
	)

	// 初始化仓库
	userRepo := repository.NewUserRepository(db)
	merchantRepo := repository.NewMerchantRepository(db)
	pageRepo := repository.NewPageRepository(db)

	// 初始化服务
	userService := service.NewUserService(userRepo)
	merchantService := service.NewMerchantService(merchantRepo)
	pageService := service.NewPageService(pageRepo)

	// 初始化控制器
	userCtrl := controller.NewUserController(userService)
	merchantCtrl := controller.NewMerchantController(merchantService)
	pageCtrl := controller.NewPageController(pageService)

	// 设置路由
	r := gin.Default()

	// 公共路由
	public := r.Group("/api")
	{
		public.POST("/register", userCtrl.Register)
		public.POST("/login", userCtrl.Login)
	}

	// 认证路由
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/user/me", userCtrl.GetCurrentUser)
		auth.POST("/merchants", merchantCtrl.CreateMerchant)

		// 页面管理路由
		pages := auth.Group("/merchants/:merchantID/pages")
		{
			pages.GET("", pageCtrl.GetPages)
			pages.POST("", pageCtrl.CreatePage)
			pages.GET("/:id", pageCtrl.GetPage)
			pages.PUT("/:id", pageCtrl.UpdatePage)
			pages.DELETE("/:id", pageCtrl.DeletePage)
		}
	}

	// 启动服务器
	log.Printf("Server started on port %s", cfg.App.Port)
	log.Fatal(r.Run(":" + cfg.App.Port))
}
