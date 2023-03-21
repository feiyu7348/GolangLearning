// author:zfy  date:2022/6/25

package main

import "fmt"

func main() {
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 是一个空map")
	}

	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "python"
	myMap2[3] = "golang"
	fmt.Println(myMap2)

	myMap3 := map[string]string{
		"one":   "php",
		"two":   "python",
		`three`: `golang`,
	}
	fmt.Println(myMap3)

	delete(myMap3, "two") // 删除
	myMap3["three"] = "R" // 修改
	fmt.Println(myMap3)

}
