// @author:zfy
// @date:2023/8/17 22:52

package main

import "fmt"

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}

func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}

func main() {
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(3)
}
