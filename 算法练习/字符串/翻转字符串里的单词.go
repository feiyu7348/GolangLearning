// author:zfy  date:2022/8/8

package main

import "strings"

func reverseWords(s string) string {
	//1.使用双指针删除冗余的空格
	slowIndex, fastIndex := 0, 0
	b := []byte(s)
	//删除头部冗余空格
	for len(b) > 0 && fastIndex < len(b) && b[fastIndex] == ' ' {
		fastIndex++
	}
	//删除单词间冗余空格
	for ; fastIndex < len(b); fastIndex++ {
		if fastIndex-1 > 0 && b[fastIndex-1] == b[fastIndex] && b[fastIndex] == ' ' {
			continue
		}
		b[slowIndex] = b[fastIndex]
		slowIndex++
	} // 循环完最后如果有空格只会有一个空格
	//删除尾部冗余空格
	if slowIndex-1 > 0 && b[slowIndex-1] == ' ' { // 此时slowIndex等于len
		b = b[:slowIndex-1]
	} else {
		b = b[:slowIndex]
	}
	//2.反转整个字符串
	reverse1(b)
	//3.反转单个单词  i单词开始位置，j单词结束位置
	i := 0
	for i < len(b) {
		j := i
		for ; j < len(b) && b[j] != ' '; j++ {
		}
		reverse1(b)
		i = j
		i++
	}
	return string(b)
}

func reverse1(b []byte) {
	left := 0
	right := len(b) - 1
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func reverseWords1(s string) string {
	b := []byte(s)

	// 移除前面、后面、中间的空格
	slow := 0
	for i := 0; i < len(b); i++ {
		if b[i] != ' ' {
			if slow != 0 { //先理解下面的for循环，再理解这段，除了第一个单词外每个单词的前面都有一个看空格
				b[slow] = ' '
				slow++
			}
			for i < len(b) && b[i] != ' ' {
				b[slow] = b[i]
				slow++
				i++
			}
		}
	}
	b = b[0:slow]

	// 翻转整个字符串
	reverse1(b)

	//翻转每个单词
	last := 0
	for i := 0; i <= len(b); i++ {
		if b[i] == ' ' || i == len(b) {
			reverse1(b[last:i])
			last = i + 1
		}
	}

	return string(b)
}

func reverseWords151(s string) string {
	ss := strings.Fields(s)
	reverse151(&ss, 0, len(ss)-1)
	return strings.Join(ss, " ")
}

func reverse151(m *[]string, i int, j int) {
	for i <= j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		i++
		j--
	}
}
