package common

import "net/http"

// APIError 定义API错误结构
type APIError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

// NewAPIError 创建新的API错误
func NewAPIError(code int, message string) *APIError {
    return &APIError{
        Code:    code,
        Message: message,
    }
}

// HTTPStatus 返回HTTP状态码
func (e *APIError) HTTPStatus() int {
    if e.Code >= 100 && e.Code < 600 {
        return e.Code
    }
    return http.StatusInternalServerError
}    