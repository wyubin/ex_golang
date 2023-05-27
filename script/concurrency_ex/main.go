package main

import (
	"fmt"
	"script/concurrency_ex/workerpool"
	"strings"
)

func div(count int) (int, error) {
	if count == 0 {
		return 0, fmt.Errorf("can not input zero")
	}
	return (count + 1) / count, nil
}

func main() {
	wp := workerpool.NewWorkerPool[int](3)
	// prepare task
	res1 := make([]int, 10)
	var task1 workerpool.Task[int] = func(ct int) error {
		_res, err := div(ct)
		if ct > 5 {
			return fmt.Errorf("result[%d]: bigger than 5", ct)
		}
		if err != nil {
			return err
		}
		fmt.Printf("rawResult[%d]:%d\n", ct, _res)
		res1[ct] = _res
		return nil
	}

	wp.Run(task1, []int{0, 1, 2, 3, 4, 5, 7}...)
	wp.Wait()
	// wp.Run(task, []int{0, 1, 2, 3, 4, 5, 7}...)
	// wp.Wait()
	fmt.Printf("Result:%+v\n", res1)
	errs := wp.Check()
	fmt.Printf("allError[%d]:%s\n", len(strings.Split(errs.Error(), "\n")), errs)

	fmt.Print("test complete\n")
}
