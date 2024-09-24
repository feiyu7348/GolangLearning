// Package models author: zfy  date: 2024/9/24
package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title   string `binding:"required"`
	Content string `binding:"required"`
	Preview string `binding:"required"`
}
