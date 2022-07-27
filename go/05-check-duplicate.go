package main

import "fmt"

func main() {
	fmt.Println("Duplicate check")
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {
	hashmap := map[int]int{}
	for i := 0; i < len(nums); i++ {
		_, found := hashmap[nums[i]]
		if found {
			return true
		} else {
			hashmap[nums[i]] = 1
		}
	}
	return false
}
