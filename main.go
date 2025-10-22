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

	// Arrays
	// // var studentsNames = [4]string{"Ahmed", "Ali, Abdullah, Mohammed"}
	// studentsGrades := [5]int{10, 20, 40, 50, 30}
	// // studentsGrades[2] = 100
	// // fmt.Println(studentsNames[0])
	// // fmt.Println(studentsNames[1:4])
	// fmt.Println(studentsGrades)

	// // Slices
	// var numbers []int
	// names := []string{"Ahmed", "Abdullah", "Ali"}
	// namesLength := len(names)
	// names = append(names, "Mohammed")
	// namesLength1 := len(names)
	// fmt.Println(numbers)
	// fmt.Println(names)
	// fmt.Println(namesLength)
	// fmt.Println(namesLength1)
	// namesCopy := make([]string, len(names))
	// copy(namesCopy, names)
	// fmt.Println("Copy:", namesCopy)

	// //For
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("rang", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
