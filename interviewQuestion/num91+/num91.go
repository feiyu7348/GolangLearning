// @author:zfy
// @date:2023/9/11 21:39

package main

type User map[string]string

func (m User) Set(key string, value string) {
	m[key] = value
}

func main() {
	m := make(User)
	m.Set("A", "One")
}
