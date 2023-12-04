package main

import "fmt"

func main() {

	input := []int{3, 2, 4}
	target := 6

	result := twoSum(input, target)
	//complexity o(N)

	fmt.Println(result)

}

func twoSum(nums []int, target int) []int {

	sumMap := map[int]int{}

	for idx, value := range nums { // N

		_, ok := sumMap[value]
		if ok {
			return []int{sumMap[value], idx}
		}

		sumMap[target-value] = idx
	}

	return nil
}
