package main

import (
	"fmt"
	"strings"
)

func demo1() {
	a := strings.Split("apple,banana,kiwi", ",")
	fmt.Println(a)
}

func demo2() {
	a := []string{"Go", "is", "awesome"}
	b := strings.Join(a, " ")
	fmt.Println(b)
}

func demo3() {
	a := " Go is cool "
	b := strings.Fields(a)
	fmt.Println(b)
}
func demo4() {
	a := "name:Armen:developer"
	b := strings.SplitN(a, ":", 2)
	fmt.Println(b)
}

func demo5() {
	a := "2025-04-19"
	b := strings.SplitAfter(a, "-")
	fmt.Println(b)
}

func main() {
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
}
