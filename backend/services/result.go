package services

// Result 统一返回结果
type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	CodeSuccess = 0
	CodeError   = 1
)

// Success 返回成功结果
func Success(data interface{}) Result {
	return Result{Code: CodeSuccess, Message: "success", Data: data}
}

// Error 返回错误结果
func Error(message string) Result {
	return Result{Code: CodeError, Message: message, Data: nil}
}
