// 控制器集成（防刷中间件）
func RateLimit() gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Minute), 5) // 每分钟5次
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "请求频率过高"})
			return
		}
		c.Next()
	}
}

// 在登录接口应用
func Login(c *gin.Context) {
	// 验证验证码
	if !captcha.Verify(c.Query("captcha_id"), c.Query("captcha")) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "验证码错误"})
		return
	}
	// 原有登录逻辑...
}