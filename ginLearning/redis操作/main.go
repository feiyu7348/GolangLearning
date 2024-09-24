// Package redis操作
// author: zfy  date: 2024/9/25
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

type User struct {
	Username string
	Password string
}

func setData(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user.Username == "" || user.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username and password cannot be empty"})
			return
		}

		if _, err := rdb.Set(user.Username, user.Password, 0).Result(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Successfully set data"})
	}
}

func getData(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Query("username")
		if rdb.Get(username).Err() != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no username provided"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success",
			"username": username,
			"password": rdb.Get(username).Val(),
		})
	}
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0,
	})
	if _, err := rdb.Ping().Result(); err != nil {
		panic(err)
	}

	r := gin.Default()
	r.POST("/set", setData(rdb))
	r.GET("/get", getData(rdb))
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server, got error: %v", err)
		return
	}
}
