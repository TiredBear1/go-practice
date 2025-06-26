package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Print("Введи любое число:")
	fmt.Scan(&num1)
	fmt.Print("Введи ещё одно любое число:")
	fmt.Scan(&num2)
	summ := num1 + num2
	fmt.Println("Результат:", summ)
}
