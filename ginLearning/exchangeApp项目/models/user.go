// Package models
// author: zfy  date: 2024/9/24
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}
