package main

import (
	"fmt"
	"time"
)

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
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }
	// for j := 0; j < 3; j++ {
	// 	fmt.Println(j)
	// }
	// for i := range 3 {
	// 	fmt.Println("rang", i)
	// }
	// for {
	// 	fmt.Println("loop")
	// 	break
	// }
	// for n := range 6 {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	// // If/Else
	// if 7%2 == 0 {
	// 	fmt.Println("7 is even")
	// } else {
	// 	fmt.Println("7 is odd")
	// }
	// if 8%4 == 0 {
	// 	fmt.Println("8 is divisible by 4")
	// }
	// if 8%2 == 0 || 7%2 == 0 {
	// 	fmt.Println("either 8 or 7 are even")
	// }
	// if num := 9; num < 0 {
	// 	fmt.Println(num, "is negative")
	// } else if num < 10 {
	// 	fmt.Println(num, "has 1 digit")
	// } else {
	// 	fmt.Println(num, "has multiple digits")
	// }

	// // Switch
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// fmt.Println(time.Now())
	// fmt.Println(time.Now().Weekday())

	switch time.Now().Weekday() {
	case time.Friday, time.Saturday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
