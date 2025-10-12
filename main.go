package main

import (
	"firstProgram/secondPackage"
	"fmt"
)

func main() {
	// print("hello")
	const name = "Abdullah"
	var age int = 20
	fmt.Println("Hello", name)
	fmt.Println("age:", age)
	SecondFile()
	secondPackage.SecondPackageFun()
	// secondPackage.secondPackageFun() // private --> error: undefined
}
