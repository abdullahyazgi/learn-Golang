package main

import (
	"firstProgram/secondpackage"
	"fmt"
)

func main() {
	// print("hello")
	const name = "Abdullah"
	var age int = 20
	fmt.Println("Hello", name)
	fmt.Println("age:", age)
	SecondFile()
	secondpackage.SecondPackageFun()
	// secondpackage.secondPackageFun() // private --> error: undefined
}
