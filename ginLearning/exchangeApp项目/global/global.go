// Package global author: zfy  date: 2024/9/24
package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	DB      *gorm.DB
	RedisDB *redis.Client
)
