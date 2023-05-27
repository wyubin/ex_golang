package main

import (
	"fmt"
	"script/concurrency_ex/workerpool"
)

func div(count int) (int, error) {
	if count == 0 {
		return 0, fmt.Errorf("Can not input zero")
	}
	return (count + 1) / count, nil
}

func main() {
	wp := workerpool.NewWorkerPool[int](5)
	// prepare task
	resAll := make([]int, 10)
	var task workerpool.Task[int] = func(ct int) error {
		_res, err := div(ct)
		if err != nil {
			return err
		}
		fmt.Printf("rawResult[%d]:%d\n", ct, _res)
		if _res > 10 {
			return fmt.Errorf("result[%d]: bigger than 10", ct)
		}
		resAll[ct] = _res
		return nil
	}

	wp.Run(task, []int{0, 1, 2, 3, 4, 5, 7}...)
	wp.Wait()
	fmt.Printf("Result:%+v\n", resAll)
	fmt.Printf("allError:%s\n", wp.Check())

	fmt.Print("test complete\n")
}
