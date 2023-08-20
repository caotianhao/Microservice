package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

type Teacher struct {
	Person  // 继承
	Subject string
}

func (p Person) GetAge() int {
	return p.Age
}

func (p Person) GetName() string {
	return p.Name
}

func (p Person) Eat() {
	fmt.Println(p.GetName() + " is eating")
}

func main() {
	p1 := Person{
		Age:  20,
		Name: "Alice",
	}
	fmt.Println(p1.GetAge())
	p1.Eat()
	t1 := Teacher{
		Person: Person{
			Age:  20,
			Name: "Bob",
		},
		Subject: "math",
	}
	fmt.Println(t1.Subject)
}
