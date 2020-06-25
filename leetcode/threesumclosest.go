package main

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	min := math.MaxFloat64
	var closest int
	for idx, num := range nums {

		left := idx + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right] + num
			diff := math.Abs(float64(sum - target))
			if diff < min {
				min = diff
				closest = sum
			}
			closest = sum
			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return target
			}
		}

	}
	return closest
}
