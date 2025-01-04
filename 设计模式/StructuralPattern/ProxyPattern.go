// Package StructuralPattern
// author: zfy  date: 2024/12/6
package StructuralPattern

import "fmt"

// 代理模式 https://developer.aliyun.com/article/1268011

// 抽象主题：文件接口
type File interface {
	Download()
}

// 真实主题：具体文件
type RealFile struct {
	filename string
}

func (f *RealFile) Download() {
	fmt.Printf("Downloading file: %s\n", f.filename)
}

// 代理：文件代理
type FileProxy struct {
	realFile *RealFile
}

func (p *FileProxy) Download() {
	p.authenticate()
	p.realFile.Download()
	p.log()
}

func (p *FileProxy) authenticate() {
	fmt.Println("Authenticating user...")
}

func (p *FileProxy) log() {
	fmt.Println("Logging download activity...")
}

// 客户端代码
func main() {
	file := &FileProxy{
		realFile: &RealFile{
			filename: "example.txt",
		},
	}
	file.Download()
}
