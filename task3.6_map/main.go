package main

import "fmt"

/*
	task 1

func main() {

	diary := make(map[string]string)

	diary["Dog"] = "собака"
	diary["Cat"] = "кошка"
	diary["Cow"] = "корова"

	var s string
	fmt.Println("Введите имя животного")
	fmt.Scan(&s)

	val, ok := diary[s]
	if ok {
		fmt.Println("В списке имеется перевод:", val)
	} else {
		fmt.Println("перевода нет")
	}

}
*/
/*
	task 2
func main() {
	diary := make(map[string]string)
	diary["Dog"] = "собака"
	diary["Cat"] = "кошка"
	diary["Cow"] = "корова"
	fmt.Println("Слова в спискеписок слов такой:", diary)

	var s string
	var s1 string
	fmt.Println("Введите имя, которое хотите удалить")
	fmt.Scan(&s)
	fmt.Println("Введите имя, которое хотите найти")
	fmt.Scan(&s1)
	delete(diary, s)
	fmt.Println(diary)
	val, ok := diary[s1]
	if ok {
		fmt.Println("В списке имеются перевод:", val)
	} else {
		fmt.Println("перевода нет")
	}
}
*/

// task 3

func main() {
	dictionary := make(map[string]int)
	var name string

	for i := 0; i < 5; i++ {
		fmt.Print("Введите имя: ")
		fmt.Scan(&name)
		dictionary[name]++
	}

	fmt.Println("Результаты:")

	for user, count := range dictionary {
		fmt.Println(user, "написал сообщений:", count, "раз")
	}
}
