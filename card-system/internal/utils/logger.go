package utils

import "log"

// Logger 简单日志工具
func Logger(msg string) {
	log.Printf("[INFO] %s", msg)
}
