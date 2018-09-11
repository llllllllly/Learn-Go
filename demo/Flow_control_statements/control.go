package main

import (
	"fmt"
)

func ifControl() {
	fmt.Println("if 条件语句")
	i := 2
	if i > 1 {
		fmt.Println("大于1")
	}
	if i := 1; i > 1 {
		fmt.Println("大于1")
	} else {
		fmt.Println("小于等于1")
	}
}

func forControl() {
	fmt.Println("for 循环语句")
	i := 0
	for {
		i++
		if i > 2 {
			break
		} 
	}
	fmt.Println("i =", i)
	for i < 6 {
		i++
	}
	fmt.Println("i =", i)
	for i := 0; i < 3; i++ {
		fmt.Println("i =", i)
	}
}

func main() {
	ifControl()
	forControl()
}