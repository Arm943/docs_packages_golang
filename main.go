package main

import (
	"fmt"
	"strings"
)

var text = "GoLang Is FUN"

func demo() {
	toUp := strings.ToUpper(text)
	toLo := strings.ToLower(text)
	fmt.Println(toUp, toLo)

}

func demo1() {
	text1 := " Go is cool "
	d := strings.TrimSpace(text1)
	fmt.Println(d)

}

func demo2() {
	text2 := "***--Go.Lang--***"
	a := strings.TrimLeft(text2, "*-.")
	b := strings.TrimRight(a, "*-.")
	fmt.Println(b)
}

func demo3() {
	text3 := "I love Go. Go is powerful. Go rocks!"
	a := strings.Replace(text3, "Go", "Golang", 1)
	b := strings.Replace(text3, "Go", "Golang", 4)
	fmt.Println(a)
	fmt.Println(b)
}

func demo4() {
	a := strings.Repeat("Go", 5)
	fmt.Println(a)
}

func main() {
	demo()
	demo1()
	demo2()
	demo3()
	demo4()
}
