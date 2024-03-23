package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 2, 4, 5, 2, 7}
	count := removeElement(a, 2)
	fmt.Println(a[:count])
}

func removeElement(nums []int, val int) int {
    var w, r int
    
    for r < len(nums) {
        if nums[r] != val {
            nums[w] = nums[r]
            w++
        }
        r++
    }
    
    return w
}
