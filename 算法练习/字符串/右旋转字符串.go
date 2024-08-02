// @author:zfy
// @date:2024/8/2 9:44

package main

func reverseRightWords(s string, target int) string {
	b := []byte(s)

	reverse2(b, 0, len(b)-1)
	reverse2(b, 0, target-1)
	reverse2(b, target, len(s)-1)
	return string(b)
}
