// author:zfy  date:2022/9/25

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	
	r.GET("/someJson", func(c *gin.Context) {
		data := map[string]interface{} {
			"lang": "GO语言",
			"tag": "<br>",
		}
		
		c.AsciiJSON(http.StatusOK, data)
	})

	err := r.Run(":8080")
	if err != nil {
		return 
	}
}
