package main

import (
	"errors"
	"fmt"
)

// ErrorCustom1, 定義 Error() 就是 一個 error type
type ErrorCustom1 struct {
	msg string
}

func (e ErrorCustom1) Error() string {
	return fmt.Sprintf("客製化 error 1:%s", e.msg)
}

type errorAnother struct{}

func (e errorAnother) Error() string {
	return "errorAnother"
}

func main() {
	// 可以用 msg 實做不一樣的 error subtype
	errCustom := ErrorCustom1{msg: "有 msg"}
	errCustom1 := fmt.Errorf("errCustom1: %w", ErrorCustom1{})
	if errors.Is(errCustom1, errCustom) {
		fmt.Println("errCustom1 是一種 errCustom")
	} else {
		fmt.Println("errCustom 不能這樣比(有 msg)")
	}
	// 直接用 error 來管理 type，但 wrap 來做狀況細節，也可以多層包裝
	errAnother := errorAnother{}
	errAnother1 := fmt.Errorf("%w: errAnother1", errorAnother{})
	errAnother2 := fmt.Errorf("%w: errAnother2", errAnother1)
	if errors.Is(errAnother1, errAnother) {
		fmt.Println("errAnother1 是一種 errAnother1 ！")
	} else {
		fmt.Println("errAnother 不能這樣比！")
	}
	if errors.Is(errAnother2, errAnother) {
		fmt.Println("errAnother2 也是一種 errAnother1 ！")
	}
	fmt.Printf("errCustom: %+v\n", errCustom)
	fmt.Printf("errCustom1: %+v\n", errCustom1)
	fmt.Printf("errAnother: %+v\n", errAnother)
	fmt.Printf("errAnother1: %+v\n", errAnother1)
	fmt.Printf("errAnother2: %+v\n", errAnother2)

}
