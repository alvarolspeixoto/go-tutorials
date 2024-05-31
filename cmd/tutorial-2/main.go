package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/* var intNum uint16 = 32767
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result) */

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(len(myString))
	fmt.Println(utf8.RuneCountInString(myString))

	var myRune rune = 'a'
	fmt.Println((myRune))

	var myBoolean bool = true
	fmt.Println(myBoolean)

	var intNum3 rune
	fmt.Println(intNum3)

	myVar := "text"
	fmt.Println(myVar)

	const pi float32 = 3.1415
	fmt.Println(pi)

}
