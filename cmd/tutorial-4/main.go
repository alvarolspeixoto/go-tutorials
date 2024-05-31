package main

import "fmt"

func main() {
	// ways to create an array
	// var intArr [3]int32 = [3]int32{1, 2, 3}
	// intArr := [3]int32{1, 2, 3}
	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	// this is not a normal array!!
	var intSlice []int32 = []int32{2, 3, 4}
	fmt.Printf("The length is %v and the capacity is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 10)
	fmt.Printf("The length is %v and the capacity is %v\n", len(intSlice), cap(intSlice))

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("The length is %v and the capacity is %v\n", len(intSlice), cap(intSlice))

	// var intSlice3 []int32 = make([]int32, 3, 8)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	myMap2 := map[string]uint8{"Álvaro": 20, "Dimitri Payet": 36}
	fmt.Println(myMap2)

	// Go aways returns a value when trying to access an index of a map
	// If you want to check if the operation was successful you can use a second value
	// that is return by the map
	age, ok := myMap2["Tata"]
	if ok {
		fmt.Printf("The age is %v\n", age)
	} else {
		fmt.Println("Invalid name")
	}

	/* fmt.Println(myMap2)
	delete(myMap2, "Álvaro")
	fmt.Println(myMap2) */

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v", i, v)
	}

	// Go doesnt have while loops.
	// You can create something like a while loop like this:
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	// or else:
	i = 0

	for {
		if i > 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

}
