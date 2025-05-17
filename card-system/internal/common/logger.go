package common

import (
    "log"
    "os"
)

// Logger 定义日志接口
type Logger interface {
    Info(msg string, args ...interface{})
    Error(msg string, args ...interface{})
    Warn(msg string, args ...interface{})
    Debug(msg string, args ...interface{})
}

// DefaultLogger 默认日志实现
type DefaultLogger struct {
    infoLogger  *log.Logger
    errorLogger *log.Logger
    warnLogger  *log.Logger
    debugLogger *log.Logger
}

// NewDefaultLogger 创建默认日志实例
func NewDefaultLogger() *DefaultLogger {
    return &DefaultLogger{
        infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
        errorLogger:  log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
        warnLogger:  log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile),
        debugLogger:  log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
    }
}

// Info 记录信息日志
func (l *DefaultLogger) Info(msg string, args ...interface{}) {
    if len(args) > 0 {
        l.infoLogger.Printf(msg, args...)
    } else {
        l.infoLogger.Println(msg)
    }
}

// Error 记录错误日志
func (l *DefaultLogger) Error(msg string, args ...interface{}) {
    if len(args) > 0 {
        l.errorLogger.Printf(msg, args...)
    } else {
        l.errorLogger.Println(msg)
    }
}

// Warn 记录警告日志
func (l *DefaultLogger) Warn(msg string, args ...interface{}) {
    if len(args) > 0 {
        l.warnLogger.Printf(msg, args...)
    } else {
        l.warnLogger.Println(msg)
    }
}

// Debug 记录调试日志
func (l *DefaultLogger) Debug(msg string, args ...interface{}) {
    if len(args) > 0 {
        l.debugLogger.Printf(msg, args...)
    } else {
        l.debugLogger.Println(msg)
    }
}    