package main

import (
	"fmt"
	"slices"
	"strings"
)

type Student struct {
	ID   int
	Name string
	Age  int
}

func main() {

	students := []Student{
		{3, "Teja", 46},
		{6, "Jai", 31},
		{2, "Bhanu", 27},
		{1, "Durga", 25},
		{4, "Jai", 32},
		{5, "Jai", 31},
	}
	fmt.Println(students)
	fmt.Println(len(students), cap(students))

	slices.SortFunc(students, func(a, b Student) int {
		return a.Age - b.Age
	})
	fmt.Println(students)
	slices.SortFunc(students, func(a, b Student) int {
		return a.ID - b.ID
	})
	fmt.Println(students)
	slices.SortFunc(students, func(a, b Student) int {
		return strings.Compare(a.Name, b.Name)
	})
	fmt.Println(students)
	slices.SortFunc(students, func(a, b Student) int {
		if nameCmp := strings.Compare(a.Name, b.Name); nameCmp != 0 {
			return nameCmp
		} else if ageCmp := a.Age - b.Age; ageCmp != 0 {
			return a.Age - b.Age
		}
		return a.ID - b.ID
	})
	fmt.Println(students)
}
