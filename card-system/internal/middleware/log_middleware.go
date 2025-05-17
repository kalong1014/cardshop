package middleware

import (
    "net/http"
    "time"

    "card-system/internal/common"
)

// LogMiddleware 日志中间件
type LogMiddleware struct {
    logger common.Logger
}

// NewLogMiddleware 创建日志中间件
func NewLogMiddleware(logger common.Logger) *LogMiddleware {
    return &LogMiddleware{
        logger: logger,
    }
}

// Handle 处理请求日志
func (m *LogMiddleware) Handle(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // 使用响应包装器捕获状态码
        lw := &loggingResponseWriter{w, http.StatusOK}
        
        next.ServeHTTP(lw, r)
        
        duration := time.Since(start)
        
        m.logger.Info("%s %s %d %s", 
            r.Method, 
            r.URL.Path, 
            lw.statusCode,
            duration,
        )
    })
}

// loggingResponseWriter 用于捕获响应状态码
type loggingResponseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
    lrw.statusCode = code
    lrw.ResponseWriter.WriteHeader(code)
}    