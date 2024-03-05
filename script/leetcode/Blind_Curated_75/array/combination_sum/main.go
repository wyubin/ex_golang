package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	checkCombination(candidates, target, 0, []int{}, &res)
	return res
}

func checkCombination(candidates []int, target int, start int, combination []int, res *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*res = append(*res, combination)
		return
	}
	for i := start; i < len(candidates); i++ {
		checkCombination(candidates, target-candidates[i], i, append(combination, candidates[i]), res)
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	k := combinationSum(candidates, target)
	fmt.Printf("k:%+v\n", k)
}
