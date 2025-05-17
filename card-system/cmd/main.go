package main

import (
	"card-system/internal/controller"
	"card-system/internal/repository"
	"card-system/internal/service"
	"card-system/internal/utils"
	"card-system/pkg/config"
	"card-system/pkg/logger"
	"card-system/pkg/redis"

	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 设置日志级别
	logger.SetLevel(cfg.Log.Level)

	// 连接数据库
	db, err := connectDB(cfg)
	if err != nil {
		logger.Fatalf("Failed to connect database: %v", err)
	}

	// 连接Redis
	redisClient, err := redis.NewRedisClient(cfg)
	if err != nil {
		logger.Fatalf("Failed to connect redis: %v", err)
	}

	// 执行数据库迁移
	if err := migrateDatabase(db); err != nil {
		logger.Fatalf("Failed to migrate database: %v", err)
	}

	// 创建依赖
	cardRepo := repository.NewCardRepository(db)
	userRepo := repository.NewUserRepository(db)
	merchantRepo := repository.NewMerchantRepository(db)

	cardGenerator := utils.NewCardGenerator(db)
	cardService := service.NewCardService(cardRepo, userRepo, merchantRepo, cardGenerator, redisClient)
	cardController := controller.NewCardController(cardService)

	// 设置路由
	router := setupRouter(cardController)

	// 启动服务器
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	// 启动服务器
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("Failed to listen and serve: %v", err)
		}
	}()

	logger.Infof("Server started on port %d", cfg.Server.Port)

	// 等待中断信号优雅关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// 最大超时时间为10秒
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}

	logger.Info("Server exiting")
}

// connectDB 连接数据库
func connectDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.DB.ConnMaxLifetime) * time.Second)

	return db, nil
}

// migrateDatabase 执行数据库迁移
func migrateDatabase(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Merchant{},
		&model.Card{},
		&model.CardLog{},
		// 添加其他需要迁移的模型
	)
}

// setupRouter 设置路由
func setupRouter(cardController *controller.CardController) *gin.Engine {
	// 在生产环境中使用默认模式
	if gin.Mode() == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// API路由
	api := r.Group("/api/v1")
	{
		cards := api.Group("/cards")
		{
			cards.POST("/", cardController.CreateCard)
			cards.GET("/:id", cardController.GetCard)
			cards.GET("/", cardController.ListCards)
			cards.PUT("/:id/activate", cardController.ActivateCard)
			cards.PUT("/:id/deactivate", cardController.DeactivateCard)
		}
	}

	return r
}
