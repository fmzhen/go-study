package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。可以返回连续的斐波纳契数。
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		c := a + b
		a = b
		b = c
		return c
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
