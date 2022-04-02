package e

// use enum in error
type ErrorCode int

const (
	IpNotAllow ErrorCode = iota
	InvalidKey
)

// errorCode 共用的失敗訊息
type ErrorResponse struct {
	Code    string
	Message string
}

var errorCodes = map[ErrorCode]ErrorResponse{
	IpNotAllow: {Code: "IP-NOT-ALLOWED", Message: "不允許的連線來源"},
	InvalidKey: {Code: "INVALID-KEY", Message: "無效的連線驗證用 Key"},
}

// enum 拿 errorCode
func (e ErrorCode) Response() ErrorResponse {
	return errorCodes[e]
}

// Error
func (e ErrorCode) Error() string {
	return e.Response().Message
}
