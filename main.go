package main

import (
	"fmt"
	"strings"
)

func demo1(a string) {
	b := strings.ReplaceAll(a, "Go", "Golang")
	fmt.Println(b)
}
func demo2() {
	a := strings.NewReplacer("&", "и", "<", "меньше", ">", "больше")
	fmt.Println(a.Replace("5 < 10 & 10 > 5"))
}

func demo3() {
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" from")
	builder.WriteString(" Builder")
	fmt.Println(builder.String())
}

func demo4() {
	text := "go is great"
	title := strings.Title(text)
	fmt.Println(title)
}

func main() {
	demo1("I love Go. Go is powerful. Go is fun.")
	demo2()
	demo3()
	demo4()
}
