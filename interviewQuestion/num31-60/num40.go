// @author:zfy
// @date:2023/6/4 14:01

package main

import "fmt"

type Param map[string]interface{}

type Show struct {
	*Param
}

func main() {
	s := new(Show)
	//s.Param["day"] = 2 存在两个问题：1.map 需要初始化才能使用；2.指针不支持索引。
	// 修复代码
	p := make(Param)
	p["day"] = 2
	s.Param = &p
	tmp := *s.Param
	fmt.Println(tmp["day"])
}
