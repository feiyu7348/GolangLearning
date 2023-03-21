// author:zfy  date:2022/6/25

package main

import "fmt"

func printArray(myArray [4]int) {
	for i, v := range myArray {
		fmt.Println("index: ", i, "value:", v)
	}
}

func main() {
	var myArray1 [10]int
	myArray2 := [10]int{1, 2, 3, 4}
	myArray3 := [4]int{1, 2, 3, 4}

	for i := 0; i < 10; i++ {
		fmt.Println(myArray1[i])
	}

	for index, value := range myArray2 {
		fmt.Println("index = ", index, "value = ", value)
	}

	fmt.Printf("myArray1 types = %T\n", myArray1)
	printArray(myArray3)
}
