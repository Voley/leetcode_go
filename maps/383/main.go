package main

import(
	"fmt"
)

func main() {
	fmt.Printf("Can you construct \"hello\" ransom note with \"go to hell\" magazine? %v", canConstruct("hello", "go to hell"))
}

func canConstruct(ransomNote string, magazine string) bool {
    
    dict := map[rune]int{}
    
    for _, v := range(magazine) {
        dict[v] = dict[v] + 1
    }
    
    for _, v := range(ransomNote) {
        if dict[v] <= 0 {
            return false
        }
        dict[v] = dict[v] - 1
    }
    
    return true
}
