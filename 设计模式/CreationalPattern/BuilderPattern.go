// Package CreationalPattern
// author: zfy  date: 2024/12/6
package CreationalPattern

import "fmt"

// 建造者模式 https://developer.aliyun.com/article/1267987

// Character 产品
type Character struct {
	profession string
	weapon     string
	equipment  string
}

// CharacterBuilder 抽象建造者
type CharacterBuilder interface {
	SetProfession()
	SetWeapon()
	SetEquipment()
	GetCharacter() Character
}

// WarriorBuilder 具体建造者：战士
type WarriorBuilder struct {
	character Character
}

func (wb *WarriorBuilder) SetProfession() {
	wb.character.profession = "战士"
}

func (wb *WarriorBuilder) SetWeapon() {
	wb.character.weapon = "剑"
}

func (wb *WarriorBuilder) SetEquipment() {
	wb.character.equipment = "盾牌"
}

func (wb *WarriorBuilder) GetCharacter() Character {
	return wb.character
}

// MageBuilder 具体建造者：法师
type MageBuilder struct {
	character Character
}

func (mb *MageBuilder) SetProfession() {
	mb.character.profession = "法师"
}

func (mb *MageBuilder) SetWeapon() {
	mb.character.weapon = "法杖"
}

func (mb *MageBuilder) SetEquipment() {
	mb.character.equipment = "法术书"
}

func (mb *MageBuilder) GetCharacter() Character {
	return mb.character
}

// Director 指挥者
type Director struct {
	builder CharacterBuilder
}

func (d *Director) Construct() Character {
	d.builder.SetProfession()
	d.builder.SetWeapon()
	d.builder.SetEquipment()
	return d.builder.GetCharacter()
}

// 示例代码
func main5() {
	director := Director{}

	// 构建战士角色
	warriorBuilder := &WarriorBuilder{}
	director.builder = warriorBuilder
	warrior := director.Construct()
	fmt.Println("战士角色：", warrior)

	// 构建法师角色
	mageBuilder := &MageBuilder{}
	director.builder = mageBuilder
	mage := director.Construct()
	fmt.Println("法师角色：", mage)
}
