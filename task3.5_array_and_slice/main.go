package main

import "fmt"

// task 1

func main() {

	numbers := []int{1, 5, 7, 10, 27, 35}
	numbers = append(numbers, 40)
	numbers = append(numbers[:1], numbers[10+1:]...)
	fmt.Println("В срезе", numbers)
	fmt.Println("длина", len(numbers))

}
