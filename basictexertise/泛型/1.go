// Package 泛型
// author: zfy  date: 2024/9/27
// https://juejin.cn/post/7088478439690371086
package main

import "fmt"

type Category struct {
	ID   int32
	Name string
	Slug string
}

type Post struct {
	ID         int32
	Categories []Category
	Title      string
	Text       string
	Slug       string
}
type cacheable interface {
	Category | Post
}
type Cache[T cacheable] struct {
	data map[string]T
}

func (c *Cache[T]) Set(key string, value T) {
	c.data[key] = value
}

func (c *Cache[T]) Get(key string) (v T) {
	if v, ok := c.data[key]; ok {
		return v
	}

	return
}

func New[T cacheable]() *Cache[T] {
	c := Cache[T]{}
	c.data = make(map[string]T)

	return &c
}

func main() {
	// create a new category
	category := Category{
		ID:   1,
		Name: "Go Generics",
		Slug: "go-generics",
	}
	// create cache for blog.Category struct
	cc := New[Category]()
	// add category to cache
	cc.Set(category.Slug, category)
	fmt.Printf("cp get:%+v\n", cc.Get(category.Slug))
	// create a new post
	post := Post{
		ID: 1,
		Categories: []Category{
			{ID: 1, Name: "Go Generics", Slug: "go-generics"},
		},
		Title: "Generics in Golang structs",
		Text:  "Here go's the text",
		Slug:  "generics-in-golang-structs",
	}
	// create cache for blog.Post struct
	cp := New[Post]()
	// add post to cache
	cp.Set(post.Slug, post)

	fmt.Printf("cp get:%+v\n", cp.Get(post.Slug))
}
