package main

import(
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Printf("Is this string: \"%v\" a palindrome? %v\n", s, isPalindrome(s))
	fmt.Printf("Is this string: \"%v\" a palindrome? %v\n", "yo", isPalindrome("yo"))
}


func isPalindrome(s string) bool {
    runes := []rune(strings.ToLower(s))
    adjusted := make([]rune, 0, len(runes))
    
    for _, v := range(runes) {
        if unicode.IsLetter(v) || unicode.IsNumber(v) {
            adjusted = append(adjusted, v)
        }   
    }
    
    sp := 0
    ep := len(adjusted) - 1
    
    if ep <= 0 {
        return true
    }
    
    for sp <= ep {
        if adjusted[sp] != adjusted[ep] {
            return false
        }
        sp++
        ep--
    }
    
    return true
}
	
