package main

import (
    "card-system/internal/common"
    "card-system/internal/controller"
    "card-system/internal/middleware"
    "card-system/internal/repository"
    "card-system/internal/service"
    "card-system/internal/utils"
    "log"
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
)

func main() {
    // 初始化配置
    initConfig()
    
    // 初始化日志
    logger := common.NewDefaultLogger()
    logger.Info("应用启动中...")
    
    // 初始化数据库
    db, err := repository.NewMySQLDB(
        viper.GetString("db.host"),
        viper.GetString("db.port"),
        viper.GetString("db.user"),
        viper.GetString("db.password"),
        viper.GetString("db.name"),
    )
    if err != nil {
        logger.Error("数据库连接失败: %v", err)
        log.Fatalf("数据库连接失败: %v", err)
    }
    defer db.Close()
    
    // 初始化Redis
    redisClient := repository.NewRedisClient(
        viper.GetString("redis.host"),
        viper.GetString("redis.port"),
        viper.GetString("redis.password"),
        viper.GetInt("redis.db"),
    )
    defer redisClient.Close()
    
    // 初始化服务
    cardRepo := repository.NewCardRepository(db)
    merchantRepo := repository.NewMerchantRepository(db)
    orderRepo := repository.NewOrderRepository(db)
    userRepo := repository.NewUserRepository(db)
    
    cardService := service.NewCardService(cardRepo, redisClient, logger)
    merchantService := service.NewMerchantService(merchantRepo, logger)
    orderService := service.NewOrderService(orderRepo, cardRepo, logger)
    userService := service.NewUserService(userRepo, logger)
    
    // 初始化工具
    cardGenerator := utils.NewCardGenerator(logger)
    paymentProcessor := utils.NewPaymentProcessor(logger)
    
    // 初始化控制器
    uploadDir := viper.GetString("upload.dir")
    fileController := controller.NewFileController(logger, uploadDir)
    cardController := controller.NewCardController(cardService, cardGenerator, logger)
    merchantController := controller.NewMerchantController(merchantService, logger)
    orderController := controller.NewOrderController(orderService, paymentProcessor, logger)
    userController := controller.NewUserController(userService, logger)
    
    // 初始化中间件
    authMiddleware := middleware.NewAuthMiddleware(userService, logger)
    logMiddleware := middleware.NewLogMiddleware(logger)
    corsMiddleware := middleware.NewCorsMiddleware()
    
    // 初始化路由
    router := setupRouter(
        fileController,
        cardController,
        merchantController,
        orderController,
        userController,
        authMiddleware,
        logMiddleware,
        corsMiddleware,
    )
    
    // 启动服务器
    port := viper.GetString("server.port")
    logger.Info("服务器启动在端口: %s", port)
    if err := http.ListenAndServe(":"+port, router); err != nil {
        logger.Error("服务器启动失败: %v", err)
        log.Fatalf("服务器启动失败: %v", err)
    }
}

// setupRouter 设置路由
func setupRouter(
    fileController *controller.FileController,
    cardController *controller.CardController,
    merchantController *controller.MerchantController,
    orderController *controller.OrderController,
    userController *controller.UserController,
    authMiddleware *middleware.AuthMiddleware,
    logMiddleware *middleware.LogMiddleware,
    corsMiddleware *middleware.CorsMiddleware,
) *gin.Engine {
    // 设置Gin模式
    if gin.Mode() == gin.ReleaseMode {
        gin.SetMode(gin.ReleaseMode)
    }
    
    // 创建路由引擎
    router := gin.Default()
    
    // 应用中间件
    router.Use(corsMiddleware.Handle)
    router.Use(logMiddleware.Handle)
    
    // 静态文件服务
    router.Static("/uploads", "./uploads")
    
    // API组
    api := router.Group("/api/v1")
    
    // 文件上传接口
    files := api.Group("/files")
    {
        files.POST("", fileController.UploadFile)
        files.GET("/:filename", fileController.DownloadFile)
        files.DELETE("/:filename", fileController.DeleteFile)
        files.GET("", fileController.ListFiles)
    }
    
    // 用户接口
    users := api.Group("/users")
    {
        users.POST("/register", userController.Register)
        users.POST("/login", userController.Login)
        users.GET("/profile", authMiddleware.Handle, userController.GetProfile)
        users.PUT("/profile", authMiddleware.Handle, userController.UpdateProfile)
    }
    
    // 商户接口
    merchants := api.Group("/merchants")
    {
        merchants.POST("", authMiddleware.Handle, merchantController.CreateMerchant)
        merchants.GET("/:id", merchantController.GetMerchant)
        merchants.GET("", merchantController.ListMerchants)
        merchants.PUT("/:id", authMiddleware.Handle, merchantController.UpdateMerchant)
        merchants.DELETE("/:id", authMiddleware.Handle, merchantController.DeleteMerchant)
    }
    
    // 卡密接口
    cards := api.Group("/cards")
    {
        cards.POST("", authMiddleware.Handle, cardController.CreateCard)
        cards.POST("/batch", authMiddleware.Handle, cardController.CreateBatchCards)
        cards.GET("/:id", cardController.GetCard)
        cards.GET("", cardController.ListCards)
        cards.PUT("/:id", authMiddleware.Handle, cardController.UpdateCard)
        cards.DELETE("/:id", authMiddleware.Handle, cardController.DeleteCard)
        cards.POST("/:id/redeem", authMiddleware.Handle, cardController.RedeemCard)
    }
    
    // 订单接口
    orders := api.Group("/orders")
    {
        orders.POST("", authMiddleware.Handle, orderController.CreateOrder)
        orders.GET("/:id", authMiddleware.Handle, orderController.GetOrder)
        orders.GET("", authMiddleware.Handle, orderController.ListOrders)
        orders.POST("/:id/pay", authMiddleware.Handle, orderController.PayOrder)
        orders.POST("/callback/:channel", orderController.PaymentCallback)
    }
    
    return router
}

// initConfig 初始化配置
func initConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    viper.AddConfigPath("./config")
    
    // 设置默认值
    viper.SetDefault("server.port", "8080")
    viper.SetDefault("db.host", "localhost")
    viper.SetDefault("db.port", "3306")
    viper.SetDefault("db.user", "root")
    viper.SetDefault("upload.dir", "./uploads")
    
    // 读取配置文件
    if err := viper.ReadInConfig(); err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            // 配置文件不存在，使用默认值
            println("配置文件未找到，使用默认配置")
        } else {
            // 配置文件存在但格式错误
            println("配置文件读取错误:", err)
            os.Exit(1)
        }
    }
}    