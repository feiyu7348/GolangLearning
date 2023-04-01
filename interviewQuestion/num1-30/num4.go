// @author:zfy
// @date:2023/4/1 10:21

package main

func main() {
	//list := new([]int)
	//list = append(list, 1)  指针不能append

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...) // 第二个参数不能直接使用 slice，需使用 … 操作符
}
