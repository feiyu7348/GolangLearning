// @author:zfy
// @date:2023/4/8 10:35

package main

func hello() []string {
	return nil
}

//func main() {
//	h := hello
//	if h == nil {
//		fmt.Println("nil")
//	} else {
//		fmt.Println("not nil")
//	}
//
//	fmt.Printf("%T", h)
//}

//not nil
//func() []string

func GetValue() interface{} {
	return "2"
}

// 类型选择
func main() {
	i := GetValue()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}
