package e

import (
	"fmt"
)

// 實做 ApiError
type ApiError struct {
	Msg string
}

func (e ApiError) Error() string {
	return fmt.Sprintf("ApiError:%s", e.Msg)
}
