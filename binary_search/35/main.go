package main

import "fmt"

func main() {
	a := []int{1,3,5,6}
	val := 5
	valb := 7
	fmt.Printf("Does value %v exist in %v? If yes at what index? %v\n", val, a, searchInsert(a, val))
	fmt.Printf("Where to insert value %v in %v? At index %v", valb, a, searchInsert(a, valb))
}

func searchInsert(nums []int, target int) int {
    low := 0
    high := len(nums) - 1
    mid := 0
    
    for low <= high {
        mid = (low + high) / 2
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }
    
    return low
}
