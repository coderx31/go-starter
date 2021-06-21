package main

import (
	"fmt"
	"math"
	"strings"
)

var score = 99.5

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func updateName(name *string) {
	*name = "Coderx"
	//return name
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	//fmt.Println("Hello Coderx!")

	// strings
	// var nameOne string = "Shenal"
	// var nameTwo = "Fernando"
	// var nameThree string // just empty string
	// fmt.Println(nameOne, nameTwo, nameThree)
	// nameThree = "Coderx"
	// fmt.Println(nameOne, nameTwo, nameThree)

	// short hand syntax for declaring variables

	// nameFour := "Owen" // this is only first time (when we are declaring), cannot use outside of a function
	// fmt.Println(nameFour)

	// ints
	// var ageOne int = 22
	// var ageTwo = 22
	// var ageThree int

	// ageThree = 22

	// fmt.Println(ageOne, ageTwo, ageThree)

	// // short hand syntax
	// ageFour := 22
	// fmt.Println(ageFour)

	// bits & memory

	// var numOne int8 = 25
	// var numTwo int8 = -128 // (-128 - 127)
	// var numThree uint8 = 25 // (0 - 255) unsigned int (cannot have negative numbers) & also can assign bit size

	// float - for float we have to assign bit size either 32 or 64

	// var scoreOne float32 = -1.5

	// float64 has highest precisions

	// scoreThree := 1.5
	// fmt.Println(scoreThree)

	// fmt module

	// Print
	// fmt.Print("Hello, ")
	// fmt.Print("World \n")
	// fmt.Print("New Line \n")

	// Println
	// fmt.Println("Hello Ninjas")
	// fmt.Println("Good bye Ninjas")

	// age := 22
	// name := "Shenal"

	// fmt.Println("My age is ", age, "and my name is ", name)

	// Printf
	//fmt.Printf("My age is %v and my name is %v \n", age, name) // this is not taking to the new line - REMEMBER
	//fmt.Println("New Line")

	// fmt.Printf("age is type of %T \n", age)
	// fmt.Printf("you scored %f points! \n", 255.55)
	// fmt.Printf("you scored %0.2f points! \n", 255.55)

	// Sprintf - (save printf)
	// var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	// fmt.Println("the saved string :", str)

	// arrays
	//var ages [3]int = [3]int{20,25,30}

	// short hand
	// var ages = [3]int{20, 25, 30}
	// names := [6]string{"Shenal", "Pamudu", "Vajith", "Pasan", "Tharindu", "Ruchira"}
	// names[0] = "Coderx"
	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// slices - (use arrays under the hood)
	// var scores = []int{100, 50, 60} // if we do not give a size when declaring an array, it will be a slice
	// scores[2] = 25
	// scores = append(scores, 85)
	// fmt.Println(scores, len(scores))

	// slice ranges
	// rangeOne := names[1:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:3]
	// fmt.Println(rangeOne, rangeTwo, rangeThree)

	// rangeOne = append(rangeOne, "Yakaaa")
	// fmt.Println(rangeOne)

	//greeting := "Hello there friends"
	// fmt.Println(strings.Contains(greeting, "Hello!"))
	// fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	// fmt.Println(strings.ToUpper(greeting))
	//fmt.Println(strings.Index(greeting, "ll"))
	//fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	//fmt.Println("the original value is ", greeting)

	// sort
	// ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	// sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 30)
	// fmt.Println(index)

	// names := []string{"Shenal", "Pamudu", "Pasan", "Vajith", "Tharindu", "Ruchira"}

	// sort.Strings(names)
	// fmt.Println(names)

	// fmt.Println(sort.SearchStrings(names, "Pamudu"))

	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is: ", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is: ", i)
	// }

	//names := []string{"Shenal", "Pamudu", "Pasan", "Vajith", "Tharindu", "Ruchira"}

	// for i := 0; i < len(names); i++ {
	// 	//fmt.Println(names[i])
	// 	names[i] = "New String"
	// }

	// for index, value := range names {
	// 	fmt.Printf("the value at index %v is %v \n", index, value)
	// }

	// for _, value := range names {
	// 	fmt.Printf("the value is %v \n", value)
	// 	value = "New String"
	// }

	//fmt.Println(names)

	// age := 25
	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("age is not less than 45")
	// }

	// names := []string{"Shenal", "Pamudu", "Pasan", "Vajith", "Tharindu", "Ruchira"}

	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("continuing at pos", index)
	// 		continue
	// 	}

	// 	if index > 2 {
	// 		fmt.Println("breaking at the position", index)
	// 		break
	// 	}

	// 	fmt.Printf("the value at pos %v is %v \n", index, value)
	// }

	// calling functions
	// sayGreeting("Shenal")
	// sayBye("Shenal")

	// cycleNames([]string{"Shenal", "Pamudu", "Vajith"}, sayGreeting)

	// a1 := circleArea(10.5)
	// a2 := circleArea(15)

	// fmt.Println(a1)
	// fmt.Println(a2)

	// fmt.Printf("area is %0.2f \n", a1)
	// fmt.Printf("area is %0.3f \n", a2)

	// fn1, sn1 := getInitials("tifa lockhart")
	// fmt.Println(fn1, sn1)

	// fn2, sn2 := getInitials("shenal fernando")
	// fmt.Println(fn2, sn2)

	// sayHello("Shenal")

	// for _, v := range points {
	// 	fmt.Println(v)
	// }

	// showScore()

	// name := "Coderx"
	// name = updateName(name)
	// fmt.Println(name)

	// menu := map[string]float64{
	// 	"pie":       5.95,
	// 	"ice cream": 3.99,
	// }

	// updateMenu(menu)
	// fmt.Println(menu)

	name := "Shenal"

	//updateName(name)

	//fmt.Println("memory address of name is ", &name)

	m := &name
	// fmt.Println("memory address", m)
	// fmt.Println("value at memory address", *m)

	fmt.Println(name)

	fmt.Println("before update the name", name)
	updateName(m)
	fmt.Println("after update the name", name)

}
