// author:zfy  date:2022/9/25

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/someJson", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status": http.StatusOK,
		})
	})
	
	r.GET("/moreJson", func(c *gin.Context) {
		var msg struct {
			Name string `json:"user"`
			Message string
			Number int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		
		c.JSON(http.StatusOK, msg)
	})
	
	r.GET("/XML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
	
	r.GET("/someprotoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label: &label,
			Reps: reps,
		}
		
		c.ProtoBuf(http.StatusOK, data)
	})

	err := r.Run(":8080")
	if err != nil {
		return 
	}
}
