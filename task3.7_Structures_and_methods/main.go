package main

import "fmt"

/* task 1-2
type User struct {
	Name string
	Age  int
}

func (u User) Greet() {
	fmt.Println("My name is ", u.Name, " and i'm ", u.Age, " old!")
}

func main() {
	var n string
	var a int
	fmt.Print("Введите Имя:")
	fmt.Scan(&n)
	fmt.Print("Введите Возраст:")
	fmt.Scan(&a)
	person := User{Name: n, Age: a}
	person.Greet()
}
*/

type User struct {
	Name string
	Age  int
}

func (u User) Greet() {
	fmt.Println("My name is ", u.Name, " and I'm ", u.Age, " old!")
}

func main() {
	var name string
	var age int
	var people []User

	for i := 0; i < 5; i++ {
		fmt.Print("Введите Имя:")
		fmt.Scan(&name)
		fmt.Print("Введите Возраст:")
		fmt.Scan(&age)
		people = append(people, User{Name: name, Age: age})
	}
	for _, user := range people {
		user.Greet()
	}

}
