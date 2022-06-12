package e

import (
	"fmt"
	"net/http"
	"runtime"
)

// StatusError 用在 類 http 的 error
type StatusError struct {
	Code   int // http code
	Err    error
	Caller string // 紀錄呼叫的位置
}

func (s StatusError) Error() string {
	return s.Err.Error()
}

func (s StatusError) Status() int {
	if s.Code == 0 {
		return 200
	}
	return s.Code
}

func newError(err error, code int) StatusError {
	pc, _, line, _ := runtime.Caller(2)
	details := runtime.FuncForPC(pc)

	return StatusError{
		Code:   code,
		Err:    err,
		Caller: fmt.Sprintf("%s#%d", details.Name(), line),
	}
}

// ErrorShowMsg Show msg to client
func ErrorShowMsg(err error) StatusError {
	return newError(err, http.StatusOK)
}

// ErrorInternal 伺服器錯誤
func ErrorInternal(err error) StatusError {
	return newError(err, http.StatusInternalServerError)
}
