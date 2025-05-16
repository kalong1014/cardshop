package main

import (
	"card-system/internal/controller"
	"card-system/internal/middleware"
	"card-system/internal/repository"
	"card-system/internal/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 数据库连接
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "user:pass@tcp(127.0.0.1:3306)/card_system?charset=utf8mb4&parseTime=True&loc=Local"
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	log.Println("Connected to Database!")

	// 自动迁移模型
	db.AutoMigrate(&repository.User{}, &repository.Merchant{}, &repository.Page{})

	// 创建仓库实例
	userRepo := repository.NewUserRepository(db)
	merchantRepo := repository.NewMerchantRepository(db)
	pageRepo := repository.NewPageRepository(db)

	// 创建服务
	userService := service.NewUserService(userRepo)
	merchantService := service.NewMerchantService(merchantRepo)
	pageService := service.NewPageService(pageRepo)

	// 创建控制器
	userController := controller.NewUserController(userService)
	merchantController := controller.NewMerchantController(merchantService)
	pageController := controller.NewPageController(pageService)

	// 创建路由
	r := gin.Default()

	// 注册API路由
	api := r.Group("/api")
	{
		// 用户认证路由
		userController.RegisterRoutes(api)

		// 商户路由 (需要认证)
		merchants := api.Group("/merchants")
		merchants.Use(middleware.AuthMiddleware())
		{
			merchantController.RegisterRoutes(merchants)
			pageController.RegisterRoutes(merchants)
		}

		// 管理路由 (需要管理员权限)
		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
		{
			// 管理员路由
		}
	}

	// 启动服务器
	log.Println("Server started on port :8080")
	r.Run(":8080")
}
