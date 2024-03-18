package main

import (
	"fmt"
	"slices"
)

func solution(intv [][]int) bool {
	slices.SortFunc(intv, func(a, b []int) int { return a[0] - b[0] })
	for i := 0; i < len(intv)-1; i++ {
		if intv[i][1] > intv[i+1][0] {
			return false
		}
	}
	return true
}

func solution2(intvs [][]int) int {
	time2room := map[int]int{}
	for _, intv := range intvs {
		time2room[intv[0]]++
		time2room[intv[1]]--
	}
	timelineRoom := [][]int{}
	for timeTmp, roomChange := range time2room {
		timelineRoom = append(timelineRoom, []int{timeTmp, roomChange})
	}
	slices.SortFunc(timelineRoom, func(a, b []int) int { return a[0] - b[0] })
	currRoom, maxRoom := 0, 0
	for _, roomchange := range timelineRoom {
		// fmt.Printf("check time[%d] with change[%d]\n", roomchange[0], roomchange[1])
		currRoom += roomchange[1]
		maxRoom = slices.Max([]int{maxRoom, currRoom})
	}
	return maxRoom
}

func main() {
	s := [][]int{{7, 10}, {2, 4}}
	k := solution2(s)
	fmt.Printf("k: %+v\n", k)
}
