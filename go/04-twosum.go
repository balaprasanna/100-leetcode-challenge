package main

import "fmt"

func main() {
	fmt.Println("Two Sum")
	fmt.Println(twoSum([]int{1, 2, 3, 7, 8}, 9))
}

func twoSum(nums []int, target int) (pairs [][]int) {

	hashmap := map[int]int{}

	for i := 0; i < len(nums); i++ {
		_, found := hashmap[nums[i]]
		if found {
			pairs = append(pairs, []int{hashmap[nums[i]], i})
		} else {
			hashmap[target-nums[i]] = i
		}
	}
	return
}
