package response

import "github.com/gin-gonic/gin"

type Response struct {
	Code    ErrorCode   `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 定义统一的错误码
type ErrorCode int

// 定义错误码常量
const (
	SuccessCode      ErrorCode = 0
	InvalidParams    ErrorCode = 1001
	Unauthorized     ErrorCode = 1002
	Forbidden        ErrorCode = 1003
	NotFound         ErrorCode = 1004
	InternalError    ErrorCode = 1005
	InvalidPassword  ErrorCode = 1006
	UserExists       ErrorCode = 1007
	DomainExists     ErrorCode = 1008
	ProviderKeyError ErrorCode = 1009
)

// 定义错误码对应的消息
var codeToMessage = map[ErrorCode]string{
	SuccessCode:      "操作成功",
	InvalidParams:    "参数错误",
	Unauthorized:     "未授权",
	Forbidden:        "禁止访问",
	NotFound:         "资源未找到",
	InternalError:    "服务器内部错误",
	InvalidPassword:  "密码错误",
	UserExists:       "用户已存在",
	DomainExists:     "域名已存在",
	ProviderKeyError: "提供商密钥错误",
}

// Message 返回错误码对应的消息
func (code ErrorCode) Message() string {
	message, ok := codeToMessage[code]
	if !ok {
		return "未知错误"
	}
	return message
}

// Success 返回成功响应
func Success(ctx *gin.Context, data interface{}) {
	if ctx == nil {
		return
	}

	resp := &Response{
		Code:    SuccessCode,
		Message: SuccessCode.Message(),
		Data:    data,
	}

	ctx.JSON(200, resp)
}

// Error 返回错误响应，使用预定义的错误码和消息
func Error(ctx *gin.Context, code ErrorCode) {
	if ctx == nil {
		return
	}

	resp := &Response{
		Code:    code,
		Message: code.Message(),
		Data:    nil,
	}

	ctx.JSON(200, resp)
}

// ErrorMessage 返回自定义错误消息的响应
func ErrorMessage(ctx *gin.Context, code ErrorCode, message string) {
	if ctx == nil {
		return
	}

	resp := &Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}

	ctx.JSON(200, resp)
}
