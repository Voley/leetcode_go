package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("In array %v looking for 6, indices that add up to 6: %v", a, twoSum(a, 6))
}

func twoSum(nums []int, target int) []int {
    dict := map[int]int{}
    for i, v := range(nums) {
        value, contains := dict[target - v]
        if contains {
            return []int{i, value}
        }
        dict[v] = i
    }
    
    return []int{0, 0}
}
