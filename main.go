package main

import "fmt"

// import (
// 	"firstProgram/secondPackage"
// 	"fmt"
// )

func main() {
	// print("hello")
	// const name = "Abdullah"
	// var age int = 20
	// fmt.Println("Hello", name)
	// fmt.Println("age:", age)
	// SecondFile()
	// secondPackage.SecondPackageFun()
	// // secondPackage.secondPackageFun() // private --> error: undefined

	// // var studentsNames = [4]string{"Ahmed", "Ali, Abdullah, Mohammed"}
	// studentsGrades := [5]int{10, 20, 40, 50, 30}
	// // studentsGrades[2] = 100

	// // fmt.Println(studentsNames[0])
	// // fmt.Println(studentsNames[1:4])

	// fmt.Println(studentsGrades)

	// Slices
	var numbers []int
	names := []string{"Ahmed", "Abdullah", "Ali"}
	namesLength := len(names)
	names = append(names, "Mohammed")
	namesLength1 := len(names)
	fmt.Println(numbers)
	fmt.Println(names)
	fmt.Println(namesLength)
	fmt.Println(namesLength1)
	namesCopy := make([]string, len(names))
	copy(namesCopy, names)
	fmt.Println("Copy:", namesCopy)

}
