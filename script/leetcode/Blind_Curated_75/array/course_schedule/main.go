package main

import "fmt"

func solution(numCourses int, prerequisites [][]int) bool {
	indegrees := make([]int, numCourses)
	prer2course := make(map[int][]int, numCourses)
	finished := []int{}
	// init
	for _, prerPair := range prerequisites {
		prer2course[prerPair[1]] = append(prer2course[prerPair[1]], prerPair[0])
		indegrees[prerPair[0]]++
	}
	// add queue with no degree
	queue := []int{}
	for idx, indegree := range indegrees {
		if indegree == 0 {
			queue = append(queue, idx)
		}
	}
	fmt.Printf("start queue: %+v\n", queue)
	fmt.Printf("prer2course: %+v\n", prer2course)
	for len(queue) > 0 {
		currIdx := queue[0]
		queue = queue[1:]
		finished = append(finished, currIdx)
		fmt.Printf("finish course: %d\n", currIdx)

		// check degree
		for _, courseIdx := range prer2course[currIdx] {
			fmt.Printf("check pair: %d->%d\n", currIdx, courseIdx)
			indegrees[courseIdx]--
			if indegrees[courseIdx] == 0 {
				queue = append(queue, courseIdx)
			}
		}
	}
	return len(finished) == numCourses
}

func main() {
	numCourses := 3
	prerequisites := [][]int{{0, 1}}
	k := solution(numCourses, prerequisites)
	fmt.Printf("k: %+v\n", k)
}
