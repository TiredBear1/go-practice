package main

import "fmt"

// task 1

func main() {

	numbers := []int{1, 5, 7, 10, 27, 35}
	numbers = append(numbers, 40)
	numbers = append(numbers[:3], numbers[4:]...)
	fmt.Println("В срезе", numbers)
	fmt.Println("длина", len(numbers))

}
