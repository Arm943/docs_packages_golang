package main

import (
	"fmt"
	"log"
	"strconv"
)

func demo1() {
	result := strconv.Itoa(123)
	fmt.Println(result)
}

func demo2() {
	result, err := strconv.Atoi("456")
	if err != nil {
		log.Println("Ошибка в конвертации: ", err)
		return
	}

	fmt.Println(result + 100)
}

func demo3() {
	a := strconv.FormatInt(31, 16)
	fmt.Println(a)
}
func demo4() {
	a, err := strconv.ParseInt("1111", 2, 16)
	if err != nil {
		log.Println("Ошибка конвертации числа", err)
		return
	}
	fmt.Println(a)
}

func demo5() {
	a := strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Println(a)
}

func main() {
	demo1()
	demo2()
	demo3()
	demo4()
	demo5()
}
