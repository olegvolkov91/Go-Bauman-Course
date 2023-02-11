package main

import "fmt"

type User struct {
	Name string
}

type University struct {
	City string
	Name string
}

type Professor struct {
	Name       string
	University University
}

func (u *University) FullInfo() {
	fmt.Printf("University name: %s and City: %s", u.Name, u.City)
}

func (u *User) changeName(name string) {
	u.Name = name
}

func (p Professor) changeName(name string) {
	p.Name = name
}

func main() {
	u := User{"Gilbert"}
	fmt.Println("Before change name", u.Name) // <- Gilbert

	u.changeName("Ivan")
	fmt.Println("After change name", u.Name) // <- Ivan

	// если не использовать указатель при создании функции, то значения исходной структуры не будут изменены
	p := Professor{
		Name: "Stanislav",
		University: University{
			City: "Kharkiv",
			Name: "KHNADU",
		},
	}
	fmt.Println("Professor name before change name", p.Name) // <- Stanislav
	p.changeName("Ivan")
	// В такой ситуации имя профессора останется неизменным
	fmt.Println("Professor name after change name", p.Name) // <- Stanislav

	// В данном примере показано как можно обращаться к методам структур если они вложены в другую струкру
	p.University.FullInfo()
}
