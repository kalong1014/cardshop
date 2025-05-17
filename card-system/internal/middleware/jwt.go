func JWTAuthAdmin(c *gin.Context) {
    claims := c.MustGet("user").(jwt.MapClaims)
    if claims["role"] != "admin" {
        c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "无管理员权限"})
        return
    }
    c.Next()
}