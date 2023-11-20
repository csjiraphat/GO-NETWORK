package main

import (
	"fmt"
)

func main() {
	var t bool = true
	var f bool

	fmt.Println("t is ", t)
	fmt.Println("f is ", f)

	var age int = 22
	var favNum float64 = 5

	var complex complex128 = 5 + 5i

	var r rune = 10 // rune is an alias for int32

	fmt.Println("age is ", age, "favNum is ", favNum)
	fmt.Println("complex128 is ", complex)
	fmt.Println("rune is ", r)

	var myName string = "Jiraphat"
	fmt.Println(myName + " is a GOD")
	fmt.Println(myName[3])
	fmt.Println("length of myName is ", len(myName))

	// Arrays variables
	var arry5 [5]float64
	arry5[0] = 90
	arry5[1] = 33
	arry5[2] = 48
	arry5[3] = 66
	arry5[4] = 59
	fmt.Println(arry5)
	fmt.Println("length of arry5 is ", len(arry5))
	fmt.Println("arry5[3] is ", arry5[3])

	arry3 := [3]float64{98, 93, 77}
	fmt.Println(arry3)

	var s []float64 = arry5[0:2] // slice of an array
	fmt.Println(s)
	type Person struct {
		name string
		age  int
	}

	var p Person
	p.name = "Jiraphat"
	p.age = 22
	fmt.Println(p)

	var x int = 5
	var xPtr *int = &x
	fmt.Println(xPtr)

}
