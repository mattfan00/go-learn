package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	for i := 0; i < 3; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
