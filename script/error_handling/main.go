package main

import (
	"errors"
	"fmt"

	"example.com/error_handling/e"
)

// ErrorCustom1, 定義 Error() 就是 一個 error type
type ErrorCustom struct {
	e.ApiError
}

type ErrorCustom1 struct {
	e.ApiError
}

func main() {
	// 可以用 msg 實做不一樣的 error subtype
	errCustom := ErrorCustom{}
	errCustom.Msg = "errCustom"
	errCustom1 := fmt.Errorf("errCustom1: %w", errCustom)
	if errors.Is(errCustom1, errCustom) {
		panic("errCustom1 是一種 errCustom")
	} else {
		fmt.Println("ErrorCustom1 跟 ErrorCustom 果然不一樣")
	}
}
