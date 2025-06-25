package main

import (
	"fmt"
)

/*	task 1-2


func main() {

	numbers := []int{1, 5, 7, 10, 27, 35}
	numbers = append(numbers, 40)
	numbers = append(numbers[:3], numbers[4:]...)
	fmt.Println("В срезе", numbers)
	fmt.Println("длина", len(numbers))

}
*/

//	task 1-2

func main() {
	var onlyOddSlice []int
	var count int

	fmt.Print("Введите количество чисел в срезе: ")
	fmt.Scan(&count)

	var value int
	for num := 0; num < count; num++ {
		fmt.Printf("Введите число №%d: ", num+1)
		_, err := fmt.Scan(&value)
		if err != nil {
			fmt.Println("Ошибка ввода. Прерывание.")
			break
		}

		if value%2 != 0 {
			onlyOddSlice = append(onlyOddSlice, value)
		}
	}

	fmt.Println("Срез (только нечётные):", onlyOddSlice)
}
