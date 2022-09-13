package main

import "fmt"

func fibo() func() int {
	a1 := -1
	a2 := 1
	var sum int
	return func() int {
		sum = a1 + a2
		a1 = a2
		a2 = sum
		return sum
	}
}

func main() {
	f := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
