package main

import (
	"fmt"
)
// 数组是值类型
func main() {
	a := [...]int{99: 1}
	b := [...]int{99: 1}
	fmt.Println(a)
	fmt.Println(a == b)
	fmt.Println(sort([5]int{5, 8, 1, 3, 6}))

	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[:]
	fmt.Println(s2)
}

func sort(a [5]int) [5]int {
	for i, l := 0, len(a); i < l; i++ {
		for j := i + 1; j < l; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	return a
}