package main

import "fmt"

func main() {
	fmt.Println("Two Sum")
	fmt.Println(twoSum([]int{1, 2, 3, 7, 8}, 9))
}

func twoSum(numbers []int, target int) (out []int) {
	i := 0
	j := len(numbers) - 1

	for i < j {
		pair_total := numbers[i] + numbers[j]
		if pair_total == target {
			out = []int{i + 1, j + 1}
			return out
		} else if pair_total > target {
			j -= 1
		} else {
			i += 1
		}
	}
	return out
}
