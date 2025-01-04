// Package StructuralPattern
// author: zfy  date: 2024/12/6
package StructuralPattern

import "fmt"

// 外观模式 https://developer.aliyun.com/article/1268006

// 子系统：文件读取
type FileReader struct {
}

func (r *FileReader) ReadFile(fileName string) string {
	// 读取文件内容
	return "File Content"
}

// 子系统：压缩算法
type Compressor struct {
}

func (c *Compressor) Compress(data string) string {
	// 使用压缩算法压缩数据
	return "Compressed Data"
}

// 子系统：文件写入
type FileWriter struct {
}

func (w *FileWriter) WriteFile(fileName string, data string) {
	// 将数据写入文件
	fmt.Printf("Write data '%s' to file '%s'\n", data, fileName)
}

// 外观：文件压缩工具
type FileCompressionFacade struct {
	reader     *FileReader
	compressor *Compressor
	writer     *FileWriter
}

func NewFileCompressionFacade() *FileCompressionFacade {
	return &FileCompressionFacade{
		reader:     &FileReader{},
		compressor: &Compressor{},
		writer:     &FileWriter{},
	}
}

func (f *FileCompressionFacade) CompressFile(fileName string) {
	fmt.Println("Compressing file:", fileName)

	// 读取文件
	data := f.reader.ReadFile(fileName)

	// 压缩数据
	compressedData := f.compressor.Compress(data)

	// 写入压缩文件
	f.writer.WriteFile(fileName+".zip", compressedData)

	fmt.Println("File compression completed.")
}

// 客户端代码
func main() {
	facade := NewFileCompressionFacade()
	facade.CompressFile("example.txt")
}
