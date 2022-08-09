// author:zfy  date:2022/8/9

package main

func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	// 1. 反转前n个字符
	// 2. 反转第n到end字符
	// 3. 反转整个字符
	reverse2(b, 0, n-1)
	reverse2(b, n, len(b)-1)
	reverse2(b, 0, len(b)-1)
	return string(b)
}

// 切片是引用传递
func reverse2(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}
