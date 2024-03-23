package main

import "fmt"

func main() {
	a := []int{5, 6, 7, 0, 0, 0}
	b := []int{9, 10, 11}
	merge(a, 3, b, 3)
	fmt.Println(a)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    p1 := m - 1
    p2 := n - 1
    w := m + n - 1
    
    for {
        if p2 < 0 || w < 0 {
            break
        }
        
        if p1 < 0 {
            for i := p2; i >= 0; i-- {
                nums1[w] = nums2[p2]
                p2--
                w--
            }
            break
        }
        
        if nums1[p1] >= nums2[p2] {
            nums1[w] = nums1[p1]
            p1--
        } else {
            nums1[w] = nums2[p2]
            p2--
        }
        w--
    }
}
