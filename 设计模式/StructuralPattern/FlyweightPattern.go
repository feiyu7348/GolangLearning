// Package StructuralPattern
// author: zfy  date: 2024/12/6
package StructuralPattern

import "fmt"

// 享元模式 https://developer.aliyun.com/article/1268008

// 享元对象：字符对象
type Character struct {
	char  rune
	font  string
	size  int
	color string
}

func (c *Character) Display() {
	fmt.Printf("Character: %c, Font: %s, Size: %d, Color: %s\n", c.char, c.font, c.size, c.color)
}

// 享元工厂：字符工厂
type CharacterFactory struct {
	characters map[rune]*Character
}

func NewCharacterFactory() *CharacterFactory {
	return &CharacterFactory{
		characters: make(map[rune]*Character),
	}
}

func (f *CharacterFactory) GetCharacter(char rune, font string, size int, color string) *Character {
	key := char
	if _, ok := f.characters[key]; !ok {
		f.characters[key] = &Character{
			char:  char,
			font:  font,
			size:  size,
			color: color,
		}
	}
	return f.characters[key]
}

// 客户端代码
func main() {
	factory := NewCharacterFactory()

	char1 := factory.GetCharacter('A', "Arial", 12, "Red")
	char1.Display()

	char2 := factory.GetCharacter('A', "Times New Roman", 14, "Blue")
	char2.Display()

	char3 := factory.GetCharacter('B', "Arial", 12, "Red")
	char3.Display()
}
