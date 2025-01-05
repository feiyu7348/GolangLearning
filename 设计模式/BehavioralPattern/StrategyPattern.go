// Package BehavioralPattern
// author: zfy  date: 2024/12/6
package BehavioralPattern

import "fmt"

// 策略模式 https://developer.aliyun.com/article/1268033

// PricingContext 环境接口
type PricingContext interface {
	SetStrategy(strategy PricingStrategy)
	CalculatePrice(originalPrice float64) float64
}

// PricingStrategy 抽象策略接口
type PricingStrategy interface {
	CalculatePrice(originalPrice float64) float64
}

// RegularUserStrategy 具体策略：普通用户策略
type RegularUserStrategy struct {
}

func (s *RegularUserStrategy) CalculatePrice(originalPrice float64) float64 {
	return originalPrice
}

// VipUserStrategy 具体策略：VIP用户策略
type VipUserStrategy struct {
}

func (s *VipUserStrategy) CalculatePrice(originalPrice float64) float64 {
	return originalPrice * 0.9
}

// SuperVipUserStrategy 具体策略：超级VIP用户策略
type SuperVipUserStrategy struct {
}

func (s *SuperVipUserStrategy) CalculatePrice(originalPrice float64) float64 {
	return originalPrice * 0.8
}

// PricingService 具体环境
type PricingService struct {
	strategy PricingStrategy
}

func (s *PricingService) SetStrategy(strategy PricingStrategy) {
	s.strategy = strategy
}

func (s *PricingService) CalculatePrice(originalPrice float64) float64 {
	return s.strategy.CalculatePrice(originalPrice)
}

// 客户端代码
func main() {
	pricingService := &PricingService{}
	pricingService.SetStrategy(&RegularUserStrategy{})

	originalPrice := 100.0
	discountedPrice := pricingService.CalculatePrice(originalPrice)
	fmt.Printf("普通用户原价：%v，折扣价：%v\n", originalPrice, discountedPrice)

	pricingService.SetStrategy(&VipUserStrategy{})
	discountedPrice = pricingService.CalculatePrice(originalPrice)
	fmt.Printf("VIP用户原价：%v，折扣价：%v\n", originalPrice, discountedPrice)

	pricingService.SetStrategy(&SuperVipUserStrategy{})
	discountedPrice = pricingService.CalculatePrice(originalPrice)
	fmt.Printf("超级VIP用户原价：%v，折扣价：%v\n", originalPrice, discountedPrice)
}
