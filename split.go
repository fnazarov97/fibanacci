package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I am learning GO programming language, and I will be programmer!"
	a := strings.Fields(s)
	var m = map[string]int{}
	for _, v := range a {
		m[v] += 1
	}
	fmt.Println(m)
}
