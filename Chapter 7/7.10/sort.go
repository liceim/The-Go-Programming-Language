/*
Exercise 7.10: Thesort.Interfacetype can be adapted to other uses. Write a functionIsPalindrome(s sort.Interface) boolthat reports whether the sequencesis a palindrome, 
in other words, reversing the sequence would not change it. Assume that the elements at indicesiandjare equal if!s.Less(i, j) && !s.Less(j, i).
*/

package main

import (
	"sort"
	"fmt"
)

// IsPalindrome report whether s is a palindrome
func IsPalindrome(s sort.Interface) bool {
	i, j := 0, s.Len()-1
	for j > i {
		// Less() only
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func main() {
	ints := []int{1, 2, 3, 2, 1}
	fmt.Println(IsPalindrome(sort.IntSlice(ints)))

	strings := []string{"hello", "world", "world", "hello"}
	fmt.Println(IsPalindrome(sort.StringSlice(strings)))
}
