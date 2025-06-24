package main

import "fmt"

func main() {
	var i int
	numbers := []int{1, 5, 7, 10, 27}
	numbers = append(numbers, 40)
	numbers = append(numbers[:i], numbers[i+1:]...)
	fmt.Println("Длина:", len(numbers))
}
