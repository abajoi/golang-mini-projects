package main

import "fmt"

type Student struct {
	name string
	age  int
}

func main() {
	stud1 := Student{}
	stud1.PrintDetails()
}

func (s Student) PrintDetails() {
	fmt.Println("Student Details\n---------------")
	fmt.Println("Name :", s.name)
	fmt.Println("Roll Number :", s.age)
}
