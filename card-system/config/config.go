func Init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv() // 优先读取环境变量

	// 数据库配置默认值（从环境变量覆盖）
	viper.SetDefault("database.dsn", os.Getenv("DB_DSN"))
	viper.SetDefault("jwt.secret", os.Getenv("JWT_SECRET"))
	viper.SetDefault("redis.addr", os.Getenv("REDIS_ADDR"))
	viper.SetDefault("storage.driver", "local") // 支持local/minio
}