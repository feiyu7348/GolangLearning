// @author:zfy
// @date:2023/5/22 22:36

package main

import (
	"fmt"
)

/*main
 * @Description: break高级用法
 */
func main() {
label1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break label1
			}
			fmt.Println("j = ", j)
		}
	}
}
