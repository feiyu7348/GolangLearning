// author:zfy  date:2022/9/25

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{} {
			"foo": "bar",
		}
		
		c.JSONP(http.StatusOK, data)
	})

	err := r.Run(":8080")
	if err != nil {
		return 
	}
}
