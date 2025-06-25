package main

import "fmt"

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
