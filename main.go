package main

import "fmt"

// "maps"

// "time"

// import (
// 	"firstProgram/secondPackage"
// 	"fmt"
// )

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 8
}

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
	// i := 2
	// fmt.Println("Write ", i, " as ")
	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// }

	// // fmt.Println(time.Now())
	// // fmt.Println(time.Now().Weekday())

	// switch time.Now().Weekday() {
	// case time.Friday, time.Saturday:
	// 	fmt.Println("It's the weekend")
	// default:
	// 	fmt.Println("It's a weekday")
	// }

	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("It's before noon")
	// default:
	// 	fmt.Println("It's after noon")
	// }

	// whatAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case bool:
	// 		fmt.Println("I'm bool")
	// 	case int:
	// 		fmt.Println("I'm an int")
	// 	default:
	// 		fmt.Printf("Don't know type %T\n", t)
	// 	}
	// }
	// whatAmI(true)
	// whatAmI(1)
	// whatAmI("hey")

	// // Maps
	// m := make(map[string]int)
	// m["k1"] = 8
	// m["k2"] = 19
	// m["w3"] = 4
	// fmt.Println("map:", m)

	// v1 := m["k1"]
	// fmt.Println("v1:", v1)

	// v3 := m["k3"]
	// fmt.Println("v3:", v3)

	// fmt.Println("len:", len(m))

	// _, prs := m["k2"]
	// fmt.Println("prs:", prs)

	// delete(m, "k2")
	// fmt.Println("map:", m)

	// clear(m)
	// fmt.Println("map:", m)

	// _, prsD := m["k2"]
	// fmt.Println("prs:", prsD)

	// n := map[string]int{"foo": 1, "bar": 2}
	// fmt.Println("map:", n)

	// n2 := map[string]int{"foo": 1, "bar": 2}

	// if maps.Equal(n, n2) {
	// 	fmt.Println("n === n2")
	// }

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

}
